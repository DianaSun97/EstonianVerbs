package initializers

import (
	"fmt"
	"os"

	"github.com/DianaSun97/PaginationScratch/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	err := DB.AutoMigrate(&models.Words{})
	if err != nil {
		return
	}
}

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to db")
	}
}
