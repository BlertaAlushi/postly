package repositories

import (
	"postly/configs"
	"postly/models"
)

type FollowRepository struct{}

func NewFollowRepository() *FollowRepository {
	return &FollowRepository{}
}

func (service *FollowRepository) Store(follow models.Follow) error {
	_, err := configs.DB.Exec("insert into follows (user_id,follow_id) values ($1, $2) on conflict (user_id, follow_id) do nothing", follow.UserID, follow.FollowID)
	return err
}

func (service *FollowRepository) Delete(follow models.Follow) error {
	_, err := configs.DB.Exec("delete from follows where user_id = $1 and follow_id = $2", follow.UserID, follow.FollowID)
	return err
}
