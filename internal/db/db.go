package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID    string `gorm:"column=id;primaryKey"`
	Value string
}

func InitDB() *gorm.DB {
	// Read DB config from env vars
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
	log.Fatalf("Missing required DB env vars. Got: user=%q, pass=****, host=%q, port=%q, name=%q", user, host, port, name)
}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	log.Fatalf("Failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&Record{}); err != nil {
	log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
