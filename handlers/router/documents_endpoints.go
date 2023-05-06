package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
	"ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		documents, err := repository.GetAllDocuments(db)
		if err != nil {
			ctx.JSON(200, documents)
			return
		}
		ctx.Status(500)
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

		if !services.IsDocumentOk(newDocument) {
			ctx.Status(400)
			return
		}
		err := repository.CreateDocument(db, &newDocument)
		if err != nil {
			ctx.Status(500)
			return
		}

	}
}
