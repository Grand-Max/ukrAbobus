package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
	"ukrabobus/models"
)

type AuthData struct {
	Email    string
	Password string
}

func AuthUser(db *gorm.DB, authData AuthData) (string, error) {
	var user models.User
	db.Find(&user, "email = ?", authData.Email)
	if user.Email != "" {
		return "", errors.New("unauthorized")
	}

	if user.Password != authData.Password {
		return "", errors.New("unauthorized")

	}

	t, _, err := CreateToken(&user)
	return t, err
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
