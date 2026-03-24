package repositories

import (
	"postly/configs"
	"postly/models"
)

type PostRepository struct{}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (pr PostRepository) Store(newPost models.Post) error {
	_, err := configs.DB.Exec("insert into posts(user_id, content) values($1,$2)", newPost.UserID, newPost.Content)
	return err
}

func (pr PostRepository) GetPost(postID int) (models.Post, error) {
	var post models.Post
	err := configs.DB.QueryRow("select * from posts where id = $1", postID).Scan(
		&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
	return post, err
}

func (pr PostRepository) Update(post models.Post) (bool, error) {
	result, err := configs.DB.Exec("update posts set content = $1 where id = $2", post.Content, post.ID)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (pr PostRepository) Delete(postID int) error {
	_, err := configs.DB.Exec("delete from posts where id = $1", postID)
	return err
}
