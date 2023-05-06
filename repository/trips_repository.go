package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateTrip(db *gorm.DB, newTrip *models.Trip) error {
	db.Create(&newTrip)
	return nil
}

func GetAllTrips(db *gorm.DB) ([]models.Trip, error) {
	var trips []models.Trip
	db.Find(&trips)
	return trips, nil
}
