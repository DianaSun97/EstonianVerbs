package initializers

import (
	"github.com/DianaSun97/PaginationScratch/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
	DB, err = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to db")
	}
}
