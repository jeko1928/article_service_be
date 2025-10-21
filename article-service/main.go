package main

import (
	"article-service/handlers" // <--- pastikan ini sesuai dengan module di go.mod
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Gunakan handler dari package handlers
	router.POST("/article", handlers.CreatePost)
	router.GET("/articles/:limit/:offset", handlers.GetPosts)
	router.GET("/article/:id", handlers.GetPostByID)
	router.PUT("/article/:id", handlers.UpdatePost)
	router.DELETE("/article/:id", handlers.DeletePost)

	router.Run(":8080")
}
