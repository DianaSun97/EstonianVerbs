package main

import (
	"log"
	"os"

	"github.com/DianaSun97/PaginationScratch/controllers"
	"github.com/DianaSun97/PaginationScratch/initializers"
	"github.com/DianaSun97/PaginationScratch/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate() {
	// Connect to SQLite database
	sqliteDB, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to SQLite database:", err)
	}

	// Connect to PostgreSQL database
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
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

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		initializers.LoadEnvVariables()
		migrate()
		return
	}

	// Initialize the app
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.CreateWords()

	// Create gin app
	app := gin.Default()

	// Configure app
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	// Routing
	app.GET("/", controllers.WordsIndexGET)
	app.POST("/", controllers.WordsIndexPOST)

	// Start app
	err := app.Run()
	if err != nil {
		log.Fatalln("Failed to start the Gin server:", err)
	}
}
