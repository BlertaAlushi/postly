package services

import (
	"net/http"
	"postly/models"
	"postly/repositories"
)

var postRepository = repositories.NewPostRepository()

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (ps PostService) GetPosts(userID int) (int, []models.Post, error) {
	return http.StatusOK, nil, nil
}

func (ps PostService) CreatePost(newPost models.Post) (int, error) {
	err := postRepository.Store(newPost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ps PostService) GetPost(postID int) (int, models.Post, error) {
	post, err := postRepository.GetPost(postID)
	if err != nil {
		return http.StatusInternalServerError, models.Post{}, err
	}
	return http.StatusOK, post, nil
}

func (ps PostService) UpdatePost(post models.Post) (int, error) {
	err := postRepository.Update(post)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ps PostService) DeletePost(postID int) (int, error) {
	err := postRepository.Delete(postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
