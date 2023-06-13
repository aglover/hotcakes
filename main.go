package main

import (
	"beacon50/tonic/controllers"
	"beacon50/tonic/initializers"
	"beacon50/tonic/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello app engine world!")
	})

	r.GET("/old/word/:id", controllers.FindWordById)
	r.GET("/word/:spelling", controllers.FindWord)
	r.POST("/words", controllers.CreateWord)

	r.GET("/doesitwork", func(c *gin.Context) {
		newWord := models.Word{
			Word:         "TEST WORD",
			PartOfSpeech: "N",
			Definitions: []models.Definition{
				{
					Definition: "this is an example",
				},
			},
			Synonyms: []models.Synonym{
				{
					Synonym: "example syn",
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		initializers.ConnectDB()
		result := initializers.DB.Create(&newWord)

		if result.Error != nil {
			c.String(http.StatusInternalServerError, "Failure!")
		} else {
			c.String(http.StatusOK, "It must have worked?")
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
