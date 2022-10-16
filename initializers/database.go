package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	DB_URL := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}
