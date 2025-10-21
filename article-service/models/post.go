package models

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_date"`
	UpdatedAt string `json:"updated_date"`
}
