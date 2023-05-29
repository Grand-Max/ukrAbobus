package service

import (
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

type UserService struct {
	repo *repos.UserRepo
}

func NewUserService(userRepo *repos.UserRepo) *UserService {
	return &UserService{
		repo: userRepo,
	}
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	users, err := service.repo.GetAllUsers()
	return users, err
}

func (service *UserService) CreateUser(newUser models.User) error {
	err := service.repo.CreateUser(&newUser)
	return err
}

func (service *UserService) IsUserOk(user models.User) bool {
	var isOk = true

	if user.FirstName == "" ||
		user.LastName == "" ||
		user.Password == "" ||
		user.Email == "" {
		isOk = false
	}

	return isOk
}
