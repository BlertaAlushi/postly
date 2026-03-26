package services

import (
	"database/sql"
	"errors"
	"net/http"
	"postly/models"
	"postly/repositories"
)

var followRepository = repositories.NewFollowRepository()

type FollowService struct{}

func NewFollowService() *FollowService {
	return &FollowService{}
}

func (service *FollowService) Follow(follow models.Follow) (int, error) {
	if follow.UserID == follow.FollowID {
		return http.StatusBadRequest, errors.New("you cannot follow yourself")
	}
	_, err := userRepository.GetUserByID(follow.FollowID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("user not found")
		}
	}
	err = followRepository.Store(follow)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (service *FollowService) Unfollow(follow models.Follow) (int, error) {
	err := followRepository.Delete(follow)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}
