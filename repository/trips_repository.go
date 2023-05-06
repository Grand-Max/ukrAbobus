package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateTrip(db *gorm.DB, newTrip *models.Trip) error {
	db.Create(&newTrip)
	return nil
}
