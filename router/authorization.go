package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
	"ukrabobus/models"
)

type AuthData struct {
	email    string
	password string
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authData AuthData

		if err := ctx.BindJSON(&authData); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		var user models.User
		db.Find(&user, "email = ?", authData.email)
		if user.Email != "" {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		if user.Password != authData.password {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		t, _, err := CreateToken(&user)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.SetCookie("auth_token", t, 60*60*72, "/", "localhost", false, true)
		ctx.Status(http.StatusOK)
	}

}

func CreateToken(user *models.User) (string, *time.Time, error) {
	exp := time.Now().Add(time.Hour * 72)
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": user.UserID,
		"exp": exp.Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("qwertyabobus"))
	if err != nil {
		return "", nil, err
	}

	return t, &exp, nil
}
