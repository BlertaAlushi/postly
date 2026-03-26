package models

type Follow struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	FollowID int `json:"follow_id"`
}
