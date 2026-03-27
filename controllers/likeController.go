package controllers

import (
	"postly/models"
	"postly/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var likeService = services.NewLikeService()

type LikeController struct{}

func NewLikeController() *LikeController {
	return new(LikeController)
}

func (lc LikeController) Likes(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	status, likes, err := likeService.GetLikes(postID)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"likes": likes,
	})

}

func (lc LikeController) NewLike(c *gin.Context) {
	var like models.Like
	userID, _ := c.Get("user_id")
	like.UserID = userID.(int)
	postID := c.Param("id")
	like.PostID, _ = strconv.Atoi(postID)

	status, err := likeService.NewLike(like)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "Post Liked",
	})
}

func (lc LikeController) RemoveLike(c *gin.Context) {
	var like models.Like
	userID, _ := c.Get("user_id")
	like.UserID = userID.(int)
	postID := c.Param("id")
	like.PostID, _ = strconv.Atoi(postID)

	status, err := likeService.RemoveLike(like)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "Like Removed",
	})
}
