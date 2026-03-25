package repositories

import (
	"postly/configs"
	"postly/models"
)

type CommentRepository struct{}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}

func (cr CommentRepository) GetComments(postId int) ([]models.Comment, error) {
	return []models.Comment{}, nil
}
func (cr CommentRepository) GetComment(comment models.Comment) (models.Comment, error) {
	err := configs.DB.QueryRow("select * from comments where id=$1 and post_id =$2", comment.ID, comment.PostID).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt)
	return comment, err
}

func (cr CommentRepository) Store(comment models.Comment) error {
	_, err := configs.DB.Exec("insert into comments(user_id, post_id,content) values($1, $2, $3)", comment.UserID, comment.PostID, comment.Content)
	return err
}

func (cr CommentRepository) Update(comment models.Comment) (bool, error) {
	result, err := configs.DB.Exec("update comments set content= $1 where user_id = $2 and post_id = $3", comment.Content, comment.UserID, comment.PostID)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	return rows > 0, err
}

func (cr CommentRepository) Delete(commentId int) error {
	_, err := configs.DB.Exec("delete from comments where id = $1", commentId)
	return err
}
