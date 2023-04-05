package router

import (
	"awesomeProject1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

		db.Create(&newUser)
		ctx.IndentedJSON(http.StatusCreated, newUser)
	}
}
