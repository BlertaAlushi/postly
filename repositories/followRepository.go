package repositories

import (
	"postly/configs"
	"postly/models"
)

type FollowRepository struct{}

func NewFollowRepository() *FollowRepository {
	return &FollowRepository{}
}

func (fr *FollowRepository) Store(follow models.Follow) error {
	_, err := configs.DB.Exec(`
		insert into follows (user_id,follow_id) values ($1, $2) 
		on conflict (user_id, follow_id) do nothing
		`, follow.UserID, follow.FollowID)
	return err
}

func (fr *FollowRepository) Delete(follow models.Follow) error {
	_, err := configs.DB.Exec(`
		delete from follows
	   	where user_id = $1 and follow_id = $2
		`, follow.UserID, follow.FollowID)
	return err
}

func (fr *FollowRepository) GetFollowing(userID int) ([]models.UserResponse, error) {
	var follows []models.UserResponse
	rows, err := configs.DB.Query(`
		select u.id, u.username, u.firstname, u.lastname
		from follows as f
		join users as u on u.id = f.follow_id
		where f.user_id = $1
		order by f.id desc
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var follow models.UserResponse
		if err = rows.Scan(
			&follow.ID,
			&follow.Username,
			&follow.Firstname,
			&follow.Lastname,
		); err != nil {
			return nil, err
		}
		follows = append(follows, follow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return follows, nil
}

func (fr *FollowRepository) GetFollowers(userID int) ([]models.UserResponse, error) {
	var followers []models.UserResponse
	rows, err := configs.DB.Query(`
		select u.id, u.username, u.firstname, u.lastname
		from follows as f
		join users as u on u.id = f.user_id
		where f.follow_id = $1
		order by f.id desc
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var follower models.UserResponse
		if err = rows.Scan(
			&follower.ID,
			&follower.Username,
			&follower.Firstname,
			&follower.Lastname,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}
