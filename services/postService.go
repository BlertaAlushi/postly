package services

import (
	"database/sql"
	"errors"
	"net/http"
	"postly/models"
	"postly/repositories"
)

var postRepository = repositories.NewPostRepository()

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (ps PostService) GetUserPosts(userID int) (int, []models.PostResponse, error) {
	_, err := userRepository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, []models.PostResponse{}, errors.New("user not found")
		}
		return http.StatusInternalServerError, nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	posts, err := postRepository.GetPosts(userID, "user")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return http.StatusOK, posts, nil
}

func (ps PostService) Feed(userID int) (int, []models.PostResponse, error) {
	posts, err := postRepository.GetPosts(userID, "feed")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return http.StatusOK, posts, nil
}

func (ps PostService) Explore(userID int) (int, []models.PostResponse, error) {
	posts, err := postRepository.GetPosts(userID, "explore")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		}
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, posts, nil
}

func (ps PostService) CreatePost(newPost models.Post) (int, error) {
	err := postRepository.Store(newPost)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}

func (ps PostService) GetPost(postID int) (int, models.Post, error) {
	post, err := postRepository.GetPost(postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, models.Post{}, errors.New("post not found")
		}
		return http.StatusInternalServerError, models.Post{}, errors.New("internal server error")
	}
	return http.StatusOK, post, nil
}

func (ps PostService) UpdatePost(post models.Post) (int, error) {
	getPost, err := postRepository.GetPost(post.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("post not found")
		}
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	if getPost.UserID != post.UserID {
		return http.StatusForbidden, errors.New("user not allowed to modify post")
	}

	_, err = postRepository.Update(post)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	return http.StatusOK, nil
}

func (ps PostService) DeletePost(post models.Post) (int, error) {
	getPost, err := postRepository.GetPost(post.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusNotFound, errors.New("post not found")
		}
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	if getPost.UserID != post.UserID {
		return http.StatusForbidden, errors.New("user not allowed to delete post")
	}
	err = postRepository.Delete(post.ID)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}
	return http.StatusOK, nil
}
