package models

import "time"

type Post struct {
	ID        int
	UserID    int
	Content   string `json:"content" binding:"required"`
	CreatedAt time.Time
}
