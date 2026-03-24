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
	userID, _ := c.Get("user_id")
	comment.UserID = userID.(int)
	postID := c.Param("id")
	comment.PostID, _ = strconv.Atoi(postID)

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
	commentId, _ := strconv.Atoi(c.Param("comment_id"))
	status, comment, err := commentService.GetComment(commentId)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"comment": comment})
}

func (cc CommentController) EditComment(c *gin.Context) {
	var comment models.Comment
	userID, _ := c.Get("user_id")
	comment.UserID = userID.(int)
	postID := c.Param("id")
	comment.PostID, _ = strconv.Atoi(postID)
	commentID := c.Param("comment_id")
	comment.ID, _ = strconv.Atoi(commentID)

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
	commentId, _ := strconv.Atoi(c.Param("comment_id"))
	status, err := commentService.DeleteComment(commentId)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": "Comment Deleted"})
}
