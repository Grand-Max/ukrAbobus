package router

import (
	"awesomeProject1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var documents []models.Document
		db.Find(&documents)
		ctx.JSON(200, documents)
	}
}

func CreateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newDocument models.Document

		if err := ctx.BindJSON(&newDocument); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		db.Create(&newDocument)
		ctx.IndentedJSON(http.StatusCreated, newDocument)
	}
}
