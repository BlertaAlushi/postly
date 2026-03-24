package services

import (
	"database/sql"
	"errors"
	"net/http"
	"postly/models"
	"postly/repositories"
)

var commentRepository = repositories.NewCommentRepository()

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (cs CommentService) NewComment(comment models.Comment) (int, error) {
	_, err := postRepository.GetPost(comment.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("post not found")
		}
		return http.StatusInternalServerError, errors.New("comment not found")
	}

	err = commentRepository.Store(comment)
	if err != nil {
		return http.StatusInternalServerError, errors.New("comment not found")
	}
	return http.StatusOK, nil
}

func (cs CommentService) GetComment(commentId int) (int, models.Comment, error) {
	comment, err := commentRepository.GetComment(commentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, models.Comment{}, errors.New("comment not found")
		}
		return http.StatusInternalServerError, models.Comment{}, errors.New("internal server error")
	}
	return http.StatusOK, comment, nil
}

func (cs CommentService) DeleteComment(commentId int) (int, error) {
	err := commentRepository.Delete(commentId)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (cs CommentService) UpdateComment(comment models.Comment) (int, error) {
	updated, err := commentRepository.Update(comment)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	if !updated {
		return http.StatusNotFound, errors.New("comment not found")
	}
	return http.StatusOK, nil
}
