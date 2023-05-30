package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(database *gorm.DB) *UserRepo {
	return &UserRepo{
		db: database,
	}
}

func (repo *UserRepo) CreateUser(newUser *models.User) error {
	repo.db.Create(&newUser)
	return nil
}

func (repo *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	repo.db.Joins("Document").Find(&users)
	return users, nil
}

func (repo *UserRepo) GetUserById(id uint) (models.User, error) {
	var user models.User
	repo.db.Find(&user, "user_id = ?", id)
	return user, nil
}

func (repo *UserRepo) DeleteUser(user *models.User) error {
	repo.db.Delete(&user)
	return nil
}
