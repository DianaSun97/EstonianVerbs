package controllers

import (
	"net/http"

	"github.com/DianaSun97/PaginationScratch/initializers"
	"github.com/DianaSun97/PaginationScratch/models"
	"github.com/gin-gonic/gin"
)

func WordsIndexGET(c *gin.Context) {
	// Get the verbs
	var words []models.Words
	initializers.DB.Find(&words)

	// Render the page
	if words == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"words": words,
		})
	}
}

func WordsIndexPOST(c *gin.Context) {
	// Get the verbs
	var words []models.Words
	initializers.DB.Find(&words)

	// Render the page
	if words == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"words": words,
		})
	}
}
