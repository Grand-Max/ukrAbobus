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

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			ctx.JSON(200, users)
			return
		}
		ctx.Status(500)
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
		if !services.IsUserOk(newUser) {
			ctx.Status(400)
			return
		}
		err := repository.CreateUser(db, &newUser)
		if err != nil {
			ctx.Status(500)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, newUser)
	}
}
