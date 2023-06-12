package controllers

import (
	"beacon50/tonic/initializers"
	"beacon50/tonic/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func FindWordById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failure!")
	} else {
		word := models.Word{}
		initializers.ConnectDB()
		initializers.DB.Preload(clause.Associations).First(&word, id)

		ctx.JSON(http.StatusOK, word)
	}
}
