package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env
	LoadEnv()

	// Get .env var
	dsn := os.Getenv("DB_URL")

	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("Error connecting to DB: %v", err)
	}
}
