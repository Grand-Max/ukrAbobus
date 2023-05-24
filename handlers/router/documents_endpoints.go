package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var docsRepo = repos.NewDocumentRepo(db)
		documents, err := docsRepo.GetAllDocuments()
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
		var docsRepo = repos.NewDocumentRepo(db)

		if err := ctx.BindJSON(&newDocument); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		if !services.IsDocumentOk(newDocument) {
			ctx.Status(400)
			return
		}
		err := docsRepo.CreateDocument(&newDocument)
		if err != nil {
			ctx.Status(500)
			return
		}

	}
}
