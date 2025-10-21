package main

import (
	"article-service/config"
	"article-service/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database
	if err := config.InitDB(); err != nil {
		log.Fatal(err)
	}

	// Buat router Gin
	r := gin.Default()

	// Tambahkan middleware CORS agar frontend bisa request
	r.Use(cors.Default())

	// Routes CRUD artikel
	r.POST("/article", handlers.CreatePost)
	r.GET("/articles/:limit/:offset", handlers.GetPosts)
	r.GET("/article/:id", handlers.GetPostByID)
	r.PUT("/article/:id", handlers.UpdatePost)
	r.DELETE("/article/:id", handlers.DeletePost)

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
