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
	api := r.Group("/api")
	{
		api.POST("/register", user.Register)
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)
		api.POST("/token/refresh", refreshToken.RefreshToken)
	}

	auth := r.Group("/api")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/posts", post.GetPosts)
		auth.POST("/posts", post.CreatePost)
		auth.GET("/posts/:id", post.GetPost)
		auth.PUT("/posts/:id", post.UpdatePost)
		auth.DELETE("/posts/:id", post.DeletePost)

		auth.POST("/posts/:id/like", like.NewLike)
		auth.DELETE("/posts/:id/like", like.RemoveLike)
	}
}
