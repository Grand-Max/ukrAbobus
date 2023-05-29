package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllDocuments(docsRepo *repos.DocumentRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		documents, err := docsRepo.GetAllDocuments()
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
