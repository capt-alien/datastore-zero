package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	ID    string `gorm:"column=id;primaryKey"`
	Value string
}

func InitDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Fatalf("Missing required DB env vars. Got: user=%q, pass=****, host=%q, port=%q, name=%q", user, host, port, name)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	var db *gorm.DB
	var err error

	// Retry connection loop
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Waiting for DB connection (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after retries: %v", err)
	}

	if err := db.AutoMigrate(&Record{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
