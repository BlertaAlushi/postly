package repositories

import (
	"postly/configs"
	"postly/models"
)

type LikeRepository struct{}

func NewLikeRepository() LikeRepository {
	return LikeRepository{}
}

func (lr LikeRepository) Store(like models.Like) error {
	_, err := configs.DB.Exec(`insert into likes(user_id, post_id) values ($1, $2) on conflict  (user_id, post_id) do nothing`,
		like.UserID, like.PostID)
	return err
}

func (lr LikeRepository) Delete(like models.Like) error {
	_, err := configs.DB.Exec(`delete from likes where user_id = $1 and post_id = $2`, like.UserID, like.PostID)
	return err
}
