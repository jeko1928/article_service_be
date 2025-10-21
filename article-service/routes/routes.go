package routes

import (
	"article-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/article", handlers.CreatePost)
	r.GET("/articles/:limit/:offset", handlers.GetPosts)
	r.GET("/article/:id", handlers.GetPostByID)
	r.PUT("/article/:id", handlers.UpdatePost)
	r.DELETE("/article/:id", handlers.DeletePost)

	return r
}
