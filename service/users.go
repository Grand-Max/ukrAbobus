package service

import "ukrabobus/models"

func IsUserOk(user models.User) bool {
	var isOk = true

	if user.FirstName == "" ||
		user.LastName == "" ||
		user.Password == "" ||
		user.Email == "" {
		isOk = false
	}

	return isOk
}
