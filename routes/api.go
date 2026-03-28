package routes

import (
	"postly/controllers"
	"postly/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	user := controllers.NewUserController()
	post := controllers.NewPostController()
	refreshToken := controllers.NewRefreshTokenController()
	like := controllers.NewLikeController()
	comment := controllers.NewCommentController()
	follow := controllers.NewFollowController()

	api := r.Group("/api")
	{
		api.POST("/register", user.Register)
		api.POST("/login", user.Login)
		api.POST("/token/refresh", refreshToken.RefreshToken)
	}

	auth := r.Group("/api")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/logout", user.Logout)
		auth.GET("/users", user.Users)

		auth.GET("/users/:id/posts", post.GetUserPosts)
		auth.GET("/feed", post.GetFeedPosts)
		auth.GET("/explore", post.GetExplorePosts)

		auth.POST("/posts", post.CreatePost)
		auth.GET("/posts/:id", post.GetPost)
		auth.PUT("/posts/:id", post.UpdatePost)
		auth.DELETE("/posts/:id", post.DeletePost)

		auth.GET("/posts/:id/likes", like.Likes)
		auth.POST("/posts/:id/like", like.NewLike)
		auth.DELETE("/posts/:id/like", like.RemoveLike)

		auth.GET("/posts/:id/comments", comment.Comments)
		auth.POST("/posts/:id/comments", comment.NewComment)
		auth.GET("/posts/:id/comments/:comment_id", comment.GetComment)
		auth.PUT("/posts/:id/comments/:comment_id", comment.EditComment)
		auth.DELETE("/posts/:id/comments/:comment_id", comment.DeleteComment)

		auth.GET("users/:id/following", follow.UserFollowing)
		auth.GET("users/:id/followers", follow.UserFollowers)

		auth.POST("/follow/:follow_id", follow.Follow)
		auth.DELETE("/follow/:follow_id", follow.Unfollow)
	}
}
