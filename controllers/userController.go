package controllers

import (
	"net/http"
	"postly/interfaces"
	"postly/models"
	"postly/services"
	"strings"

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

	interfaces.NormalizeInput(&user)

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

	interfaces.NormalizeInput(&login)

	status, token, err := userService.Login(login)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"token": token})
}

func (u UserController) Logout(c *gin.Context) {
	refreshToken := struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetInt("user_id")
	status, err := userService.Logout(userID, refreshToken.RefreshToken)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"message": "User successfully logged out."})
}

func (u UserController) Users(c *gin.Context) {
	search := struct {
		Search string `json:"search" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, users, err := userService.GetUsers(strings.TrimSpace(search.Search))
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"users": users})
}
