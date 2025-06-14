package db
import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	Key string `gorm:"primaryKey"`
	Value string
}

func InitDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate the Record struct into a real table
	if err := db.AutoMigrate(&Record{}); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}
