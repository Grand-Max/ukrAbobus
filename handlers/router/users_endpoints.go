package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllUsers(usersRepo *repos.UserRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		users, err := usersRepo.GetAllUsers()
		if err != nil {
			ctx.JSON(200, users)
			return
		}
		ctx.Status(500)
	}
}

func CreateUser(usersRepo *repos.UserRepo) gin.HandlerFunc {
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
		err := usersRepo.CreateUser(&newUser)
		if err != nil {
			ctx.Status(500)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, newUser)
	}
}
