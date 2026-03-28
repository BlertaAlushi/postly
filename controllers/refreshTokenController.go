package controllers

import (
	"net/http"
	"postly/services"

	"github.com/gin-gonic/gin"
)

var refreshTokenService = services.NewRefreshTokenService()

type RefreshTokenController struct{}

func NewRefreshTokenController() *RefreshTokenController {
	return &RefreshTokenController{}
}

func (rtc RefreshTokenController) RefreshToken(c *gin.Context) {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, token, err := refreshTokenService.RefreshToken(refreshToken.RefreshToken)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"token": token})

}
