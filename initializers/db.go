package initializers

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"
	_ "os"

	"github.com/DianaSun97/PaginationScratch/models"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	err := DB.AutoMigrate(&models.Words{})
	if err != nil {
		log.Fatalln("Failed to migrate database:", err)
	}
}

func ConnectToDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to PostgreSQL database:", err)
	}
}
