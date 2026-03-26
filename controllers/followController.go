package controllers

import (
	"postly/models"
	"postly/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var followService = services.NewFollowService()

type FollowController struct{}

func NewFollowController() *FollowController {
	return &FollowController{}
}

func (fc FollowController) Follow(c *gin.Context) {
	var follow models.Follow
	follow.UserID = c.GetInt("user_id")
	follow.FollowID, _ = strconv.Atoi(c.Param("follow_id"))

	status, err := followService.Follow(follow)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "Following User",
	})

}
func (fc FollowController) Unfollow(c *gin.Context) {
	var follow models.Follow
	follow.UserID = c.GetInt("user_id")
	follow.FollowID, _ = strconv.Atoi(c.Param("follow_id"))

	status, err := followService.Unfollow(follow)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "Unfollowing User",
	})
}

func (fc FollowController) Following(c *gin.Context) {
	userId := c.GetInt("user_id")
	status, following, err := followService.Following(userId)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"following": following,
	})
}

func (fc FollowController) Followers(c *gin.Context) {
	userId := c.GetInt("user_id")
	status, followers, err := followService.Followers(userId)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"followers": followers,
	})
}
