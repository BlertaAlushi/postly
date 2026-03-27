package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type UserComment struct {
	UserResponse
	Comment string `json:"comment"`
}
