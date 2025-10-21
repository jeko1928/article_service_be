package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Struktur data artikel
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_date"`
	UpdatedAt time.Time `json:"updated_date"`
}

// ✅ Simulasi database (sementara pakai slice, nanti bisa diganti MySQL)
var posts []Post
var nextID = 1

// 🟢 CREATE
func CreatePost(c *gin.Context) {
	var newPost Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validasi
	if len(newPost.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title minimal 20 karakter"})
		return
	}
	if len(newPost.Content) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content minimal 200 karakter"})
		return
	}
	if len(newPost.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category minimal 3 karakter"})
		return
	}
	if newPost.Status != "publish" && newPost.Status != "draft" && newPost.Status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus publish/draft/thrash"})
		return
	}

	newPost.ID = nextID
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	nextID++

	posts = append(posts, newPost)
	c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dibuat", "data": newPost})
}

// 🟡 READ ALL (dengan paging)
func GetPosts(c *gin.Context) {
	limitParam := c.Param("limit")
	offsetParam := c.Param("offset")

	limit, err1 := strconv.Atoi(limitParam)
	offset, err2 := strconv.Atoi(offsetParam)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit/Offset harus angka"})
		return
	}

	// Paging manual
	start := offset
	end := offset + limit
	if start > len(posts) {
		c.JSON(http.StatusOK, []Post{})
		return
	}
	if end > len(posts) {
		end = len(posts)
	}

	c.JSON(http.StatusOK, posts[start:end])
}

// 🟢 READ ONE
func GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	for _, p := range posts {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
}

// 🟠 UPDATE
func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedPost Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, p := range posts {
		if p.ID == id {
			updatedPost.ID = p.ID
			updatedPost.CreatedAt = p.CreatedAt
			updatedPost.UpdatedAt = time.Now()
			posts[i] = updatedPost
			c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil diperbarui", "data": updatedPost})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
}

// 🔴 DELETE
func DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Artikel berhasil dihapus"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
}
