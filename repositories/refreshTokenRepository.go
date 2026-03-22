package repositories

import (
	"database/sql"
	"errors"
	"postly/configs"
	"postly/models"
)

type RefreshTokenRepository struct{}

func NewRefreshTokenRepository() *RefreshTokenRepository {
	return &RefreshTokenRepository{}
}

func (r RefreshTokenRepository) CreateRefreshToken(refreshToken models.RefreshToken) (int, error) {
	var id int
	err := configs.DB.QueryRow("Insert into refresh_tokens(token_hash,user_id,expires_at) values($1,$2,$3) returning id",
		refreshToken.TokenHash, refreshToken.UserID, refreshToken.ExpiresAt).Scan(&id)
	return id, err
}

func (r RefreshTokenRepository) GetRefreshTokenById(tokenID int) (models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := configs.DB.QueryRow("select token_hash, user_id, expires_at from refresh_tokens where id = $1", tokenID).Scan(
		&refreshToken.TokenHash,
		&refreshToken.UserID,
		&refreshToken.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.RefreshToken{}, errors.New("refresh token not found")
		}
		return models.RefreshToken{}, err
	}
	return refreshToken, nil
}

func (r RefreshTokenRepository) DeleteRefreshToken(tokenHash string) error {
	_, err := configs.DB.Exec("delete from refresh_tokens where token_hash = $1", tokenHash)
	return err
}
