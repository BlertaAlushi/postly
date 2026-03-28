package controllers

import (
	"fmt"
	"net/http"
	"postly/models"
	"postly/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var postService = services.NewPostService()

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (p PostController) GetUserPosts(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	status, posts, err := postService.GetUserPosts(userID)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"posts": posts,
	})
}

func (p PostController) GetFeedPosts(c *gin.Context) {
	userID := c.GetInt("user_id")
	fmt.Println(userID)
	status, posts, err := postService.Feed(userID)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"posts": posts,
	})
}
func (p PostController) GetExplorePosts(c *gin.Context) {
	userID := c.GetInt("user_id")
	fmt.Println(userID)
	status, posts, err := postService.Explore(userID)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"posts": posts,
	})
}

func (p PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userID := c.GetInt("user_id")
	post.UserID = userID

	status, err := postService.CreatePost(post)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post Created",
	})
}

func (p PostController) GetPost(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, _ := strconv.Atoi(postID)
	status, post, err := postService.GetPost(postIDInt)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"post": post,
	})
}

func (p PostController) UpdatePost(c *gin.Context) {
	var post models.Post
	post.ID, _ = strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	post.UserID = userID

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	status, err := postService.UpdatePost(post)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post Edited",
	})
}

func (p PostController) DeletePost(c *gin.Context) {
	var post models.Post
	post.ID, _ = strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")
	post.UserID = userID
	status, err := postService.DeletePost(post)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "Post Deleted",
	})
}
