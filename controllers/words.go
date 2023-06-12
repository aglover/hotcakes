package controllers

import (
	"beacon50/tonic/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindWordById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failure!")
	} else {
		word := models.Word{}
		word.FindWordById(id)
		ctx.JSON(http.StatusOK, word)
	}
}
