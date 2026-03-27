package services

import (
	"database/sql"
	"errors"
	"net/http"
	"postly/models"
	"postly/repositories"
)

var likeRepository = repositories.NewLikeRepository()

type LikeService struct{}

func NewLikeService() *LikeService {
	return &LikeService{}
}

func (ls LikeService) NewLike(like models.Like) (int, error) {
	_, err := postRepository.GetPost(like.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("post not found")
		}
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	err = likeRepository.Store(like)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}

func (ls LikeService) RemoveLike(like models.Like) (int, error) {
	err := likeRepository.Delete(like)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (ls LikeService) GetLikes(postID int) (int, []models.UserResponse, error) {
	_, err := postRepository.GetPost(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, nil, errors.New("post not found")
		}
		return http.StatusInternalServerError, nil, err
	}

	likes, err := likeRepository.GetPostLikes(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, likes, nil
}
