package controllers

import (
	"aglover/hotcakes/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWord(ctx *gin.Context) {
	var newWord models.Word
	if err := ctx.BindJSON(&newWord); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error with JSON binding": err.Error()})
	}
	if errCreated := newWord.Create(); errCreated != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error creating word record": errCreated.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "word created"})
}

func FindWord(ctx *gin.Context) {
	word := models.Word{}
	word.FindWordBySpelling(ctx.Param("spelling"))
	ctx.JSON(http.StatusOK, word)
}

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
