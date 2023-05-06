package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
	services "ukrabobus/service"
)

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []models.User
		db.Find(&users)
		ctx.JSON(200, users)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User

		if err := ctx.BindJSON(&newUser); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		if services.IsUserOk(newUser) {
			db.Create(&newUser)
			ctx.IndentedJSON(http.StatusCreated, newUser)
		} else {
			ctx.Status(400)
		}

	}
}
