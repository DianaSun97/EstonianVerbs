package main

import (
	"log"

	"github.com/DianaSun97/PaginationScratch/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Connect to SQLite database
	sqliteDB, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to SQLite database:", err)
	}

	// Connect to PostgreSQL database
	dsn := "host=localhost user=your_postgres_user password=your_postgres_password dbname=your_postgres_dbname port=5432 sslmode=disable"
	postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to PostgreSQL database:", err)
	}

	// Migrate the schema
	err = postgresDB.AutoMigrate(&models.Words{})
	if err != nil {
		log.Fatalln("Failed to migrate PostgreSQL schema:", err)
	}

	// Read data from SQLite
	var words []models.Words
	result := sqliteDB.Find(&words)
	if result.Error != nil {
		log.Fatalln("Failed to read data from SQLite database:", result.Error)
	}

	// Write data to PostgreSQL
	for _, word := range words {
		postgresDB.Create(&word)
	}

	log.Println("Data migration completed successfully!")
}
