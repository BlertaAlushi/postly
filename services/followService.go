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

func (fs *FollowService) Follow(follow models.Follow) (int, error) {
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

func (fs *FollowService) Unfollow(follow models.Follow) (int, error) {
	err := followRepository.Delete(follow)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (fs *FollowService) Following(userID int) (int, []models.UserResponse, error) {
	followingUsers, err := followRepository.GetFollowing(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}
	return http.StatusOK, followingUsers, nil
}

func (fs *FollowService) Followers(userID int) (int, []models.UserResponse, error) {
	followers, err := followRepository.GetFollowers(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, errors.New("internal server error")
	}
	return http.StatusOK, followers, nil
}
