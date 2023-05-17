package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	services "ukrabobus/service"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authData services.AuthData

		if err := ctx.BindJSON(&authData); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		token, err := services.AuthUser(db, authData)

		if err == nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.SetCookie("auth_token", token, 60*60*72, "/", "localhost", false, true)
		ctx.Status(http.StatusOK)
	}

}
