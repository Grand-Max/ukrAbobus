package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateUser(db *gorm.DB, newUser *models.User) error {
	db.Create(&newUser)
	return nil
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	db.Joins("Document").Find(&users)
	return users, nil
}
