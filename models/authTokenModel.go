package models

import "time"

type AuthToken struct {
	AccessToken  string
	RefreshToken string
}

type RefreshToken struct {
	ID        int
	TokenHash string
	UserID    int
	ExpiresAt time.Time
}
