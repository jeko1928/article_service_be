package handlers

import (
	"article-service/config"
	"article-service/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 🟢 CREATE
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validasi input
	if len(newPost.Title) < 20 || len(newPost.Content) < 200 || len(newPost.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validasi gagal"})
		return
	}

	if newPost.Status != "publish" && newPost.Status != "draft" && newPost.Status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus publish/draft/thrash"})
		return
	}

	_, err := config.DB.Exec(`
        INSERT INTO posts (title, content, category, status, created_date, updated_date)
        VALUES (?, ?, ?, ?, NOW(), NOW())`,
		newPost.Title, newPost.Content, newPost.Category, newPost.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dibuat"})
}

// 🟡 READ ALL
func GetPosts(c *gin.Context) {
	limitParam := c.Param("limit")
	offsetParam := c.Param("offset")
	limit, err1 := strconv.Atoi(limitParam)
	offset, err2 := strconv.Atoi(offsetParam)
	if err1 != nil || err2 != nil || limit < 1 || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit harus >0 dan Offset >=0"})
		return
	}

	rows, err := config.DB.Query(`
		SELECT id, title, content, category, status, created_date, updated_date
		FROM posts
		ORDER BY id DESC
		LIMIT ? OFFSET ?`, limit, offset)
	if err != nil {
		log.Println("DB Query Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan di server"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Status, &p.CreatedAt, &p.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, p)
	}

	c.JSON(http.StatusOK, posts)
}

// 🟢 READ ONE
func GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var p models.Post
	err = config.DB.QueryRow(`
		SELECT id, title, content, category, status, created_date, updated_date
		FROM posts WHERE id = ?`, id).Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Status, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, p)
}

// 🟠 UPDATE
func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE posts SET title = ?, content = ?, category = ?, status = ?, updated_date = NOW()
		WHERE id = ?`,
		updatedPost.Title, updatedPost.Content, updatedPost.Category, updatedPost.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil diperbarui"})
}

// 🔴 DELETE
func DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	_, err = config.DB.Exec(`DELETE FROM posts WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dihapus"})
}
