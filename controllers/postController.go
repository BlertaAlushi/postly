package controllers

import (
	"net/http"
	"postly/models"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (p PostController) GetPosts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Unauthorized",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
	})
}
func (p PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"Post": post.Title,
	})
}
