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

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRepo = repos.NewUserRepo(db)

		users, err := userRepo.GetAllUsers()
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
		var userRepo = repos.NewUserRepo(db)

		if err := ctx.BindJSON(&newUser); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		if !services.IsUserOk(newUser) {
			ctx.Status(400)
			return
		}
		err := userRepo.CreateUser(&newUser)
		if err != nil {
			ctx.Status(500)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, newUser)
	}
}
