package db
import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
// Record represents a key/value pair in the datastore
type Record struct {
	Key string `gorm:"primaryKey"`
	Value string
}

// InitDB opens (or creates) a SQLite database at the given path and auto-migrates the Record schema
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
