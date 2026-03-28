package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type PostResponse struct {
	ID        int          `json:"id"`
	User      UserResponse `json:"user"`
	Content   string       `json:"content"`
	Likes     int          `json:"likes"`
	Comments  int          `json:"comments"`
	CreatedAt time.Time    `json:"created_at"`
}
