package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ukrabobus/models"
	services "ukrabobus/service"
)

func GetAllDocuments(service *services.DocumentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		documents, err := service.GetAllDocuments()
		if err != nil {
			ctx.JSON(200, documents)
			return
		}
		ctx.Status(500)
	}
}

func CreateDocument(service *services.DocumentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newDocument models.Document

		if err := ctx.BindJSON(&newDocument); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		if !service.IsDocumentOk(newDocument) {
			ctx.Status(400)
			return
		}
		err := service.CreateDocument(newDocument)
		if err != nil {
			ctx.Status(500)
			return
		}

	}
}
