package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"postly/models"
	"postly/repositories"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var refreshTokenRepository = repositories.NewRefreshTokenRepository()

type RefreshTokenService struct{}

func NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{}
}

func (rts RefreshTokenService) NewRefreshToken(userID int) (int, string, error) {
	return NewRT(userID)
}

func (rts RefreshTokenService) RefreshToken(postRefreshToken string) (int, models.AuthToken, error) {
	var authToken = models.AuthToken{}
	parts := strings.Split(postRefreshToken, ".")
	tokenID, _ := strconv.Atoi(parts[0])
	getRefreshToken, err := refreshTokenRepository.GetRefreshTokenById(tokenID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, authToken, errors.New("refresh token not found")
		}
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(getRefreshToken.TokenHash), []byte(parts[1]))
	if err != nil {
		return http.StatusNotFound, authToken, errors.New("wrong refresh token")
	}

	userID := getRefreshToken.UserID
	err = refreshTokenRepository.DeleteRefreshToken(getRefreshToken.TokenHash)
	if err != nil {
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	accessToken, err := GenerateJWT(userID)
	if err != nil {
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	id, newRefreshToken, err := NewRT(userID)
	if err != nil {
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	authToken.AccessToken = accessToken
	authToken.RefreshToken = fmt.Sprintf("%d.%s", id, newRefreshToken)
	return http.StatusOK, authToken, nil
}

func NewRT(userID int) (int, string, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0, "", err
	}
	refreshToken := base64.RawURLEncoding.EncodeToString(b)
	hashedToken, _ := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	newRefreshToken := models.RefreshToken{
		TokenHash: string(hashedToken),
		UserID:    userID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	id, err := refreshTokenRepository.CreateRefreshToken(newRefreshToken)
	if err != nil {
		return 0, "", err
	}
	return id, refreshToken, nil
}
