package controllers

import (
	"net/http"
	"postly/models"
	"postly/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentService = services.NewCommentService()

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (cc CommentController) Comments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (cc CommentController) NewComment(c *gin.Context) {
	var comment models.Comment
	userID := c.GetInt("user_id")
	comment.UserID = userID
	comment.PostID, _ = strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := commentService.NewComment(comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"message": "Comment Added"})
}

func (cc CommentController) GetComment(c *gin.Context) {
	var comment models.Comment
	comment.ID, _ = strconv.Atoi(c.Param("comment_id"))
	comment.PostID, _ = strconv.Atoi(c.Param("id"))
	status, comment, err := commentService.GetComment(comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"comment": comment})
}

func (cc CommentController) EditComment(c *gin.Context) {
	var comment models.Comment
	userID := c.GetInt("user_id")
	comment.UserID = userID
	comment.PostID, _ = strconv.Atoi(c.Param("id"))
	comment.ID, _ = strconv.Atoi(c.Param("comment_id"))

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := commentService.UpdateComment(comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": "Comment Updated"})
}

func (cc CommentController) DeleteComment(c *gin.Context) {
	var comment models.Comment
	userID := c.GetInt("user_id")
	comment.UserID = userID
	comment.PostID, _ = strconv.Atoi(c.Param("id"))
	comment.ID, _ = strconv.Atoi(c.Param("comment_id"))

	status, err := commentService.DeleteComment(comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": "Comment Deleted"})
}
