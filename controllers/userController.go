package controllers

import (
	"net/http"
	"postly/models"
	"postly/services"

	"github.com/gin-gonic/gin"
)

var userService = services.NewUserService()

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u UserController) Register(c *gin.Context) {
	var user models.Register
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, message, err := userService.Register(user)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": message})
}

func (u UserController) Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, token, err := userService.Login(login)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"token": token})
}

func (u UserController) Logout(c *gin.Context) {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token"`
	}{}
	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := userService.Logout(refreshToken.RefreshToken)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"message": "User successfully logged out."})
}
