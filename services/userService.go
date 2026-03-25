package services

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"postly/models"
	"postly/repositories"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var userRepository = repositories.NewUserRepository()
var refreshTokenService = NewRefreshTokenService()

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us UserService) Register(register models.Register) (int, string, error) {
	user, err := userRepository.GetUserByUsernameOrEmail(register.Username, register.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

		} else {
			return http.StatusInternalServerError, "", errors.New("internal server error")
		}
	}

	if user.Username == register.Username {
		return http.StatusConflict, "", errors.New("username already exists")
	}
	if user.Email == register.Email {
		return http.StatusConflict, "", errors.New("email already taken")
	}

	if user.Email == register.Email {
		return http.StatusConflict, "", errors.New("email already taken")
	}

	err = userRepository.CreateUser(register)
	if err != nil {
		return http.StatusInternalServerError, "", errors.New("internal server error")
	}

	return http.StatusOK, "User Created", nil
}

func (us UserService) Login(login models.Login) (int, models.AuthToken, error) {
	authToken := models.AuthToken{}
	user, err := userRepository.GetUser(login.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, authToken, errors.New("user not found")
		}
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return http.StatusUnauthorized, authToken, errors.New("wrong password")
	}

	accessToken, err := GenerateJWT(user.ID)
	if err != nil {
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	id, refreshToken, err := refreshTokenService.NewRefreshToken(user.ID)
	if err != nil {
		return http.StatusInternalServerError, authToken, errors.New("internal server error")
	}

	authToken.AccessToken = accessToken
	authToken.RefreshToken = fmt.Sprintf("%d.%s", id, refreshToken)

	return http.StatusOK, authToken, nil
}

func (us UserService) Logout(refreshToken string) (int, error) {
	parts := strings.Split(refreshToken, ".")
	tokenId, _ := strconv.Atoi(parts[0])
	getRefreshToken, err := refreshTokenRepository.GetRefreshTokenById(tokenId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusUnauthorized, errors.New("refresh token not found")
		}
		return http.StatusInternalServerError, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(getRefreshToken.TokenHash), []byte(parts[1])); err != nil {
		return http.StatusUnauthorized, errors.New("wrong refresh token")
	}

	err = refreshTokenRepository.DeleteRefreshToken(getRefreshToken.TokenHash)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}
