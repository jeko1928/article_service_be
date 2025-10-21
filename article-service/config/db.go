package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	username := "root"
	password := "Ageni2135"
	host := "127.0.0.1"
	port := "3306"
	database := "article"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, host, port, database)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi DB: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("gagal koneksi ke DB: %w", err)
	}

	fmt.Println("Berhasil koneksi ke database!")
	return nil
}
