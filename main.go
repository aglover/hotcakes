package main

import (
	"aglover/hotcakes/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/word/:id", controllers.FindWordById)
	r.GET("/words/:spelling", controllers.FindWord)
	r.POST("/words", controllers.CreateWord)

	r.Run()
}
