package main

import (
	"github.com/DianaSun97/PaginationScratch/controllers"
	"github.com/DianaSun97/PaginationScratch/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.CreateWords()
}

func main() {
	// Create gin app
	app := gin.Default()

	// Configure app
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	// Routing
	app.GET("/", controllers.WordsIndexGET)

	// Start app
	app.Run()
}
