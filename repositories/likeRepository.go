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

func (lr LikeRepository) GetPostLikes(postID int) ([]models.UserResponse, error) {
	var likes []models.UserResponse
	rows, err := configs.DB.Query(`
		select u.id, u.username, u.firstname, u.lastname
		from likes as l
		join posts as p on l.post_id = p.id
		join users as u on l.user_id = u.id
		where l.post_id = $1
		order by l.id desc
	`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var like models.UserResponse
		if err = rows.Scan(
			&like.ID,
			&like.Username,
			&like.Firstname,
			&like.Lastname,
		); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return likes, nil
}
