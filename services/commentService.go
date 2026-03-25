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
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	err = commentRepository.Store(comment)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (cs CommentService) GetComment(comment models.Comment) (int, models.Comment, error) {
	comment, err := commentRepository.GetComment(comment)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, models.Comment{}, errors.New("comment not found")
		}
		return http.StatusInternalServerError, models.Comment{}, errors.New("internal server error")
	}
	return http.StatusOK, comment, nil
}

func (cs CommentService) DeleteComment(comment models.Comment) (int, error) {
	getComment, err := commentRepository.GetComment(comment)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("comment not found")
		}
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	if getComment.UserID != comment.UserID {
		return http.StatusForbidden, errors.New("user not allowed to delete comment")
	}
	err = commentRepository.Delete(comment.ID)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (cs CommentService) UpdateComment(comment models.Comment) (int, error) {
	getComment, err := commentRepository.GetComment(comment)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("comment not found")
		}
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	if getComment.UserID != comment.UserID {
		return http.StatusForbidden, errors.New("user not allowed to modify comment")
	}

	_, err = commentRepository.Update(comment)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}
