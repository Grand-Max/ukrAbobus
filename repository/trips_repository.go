package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

type TripRepo struct {
	db *gorm.DB
}

func NewTripRepo(database *gorm.DB) *TripRepo {
	return &TripRepo{
		db: database,
	}
}

func (repo *TripRepo) CreateTrip(newTrip *models.Trip) error {
	repo.db.Create(&newTrip)
	return nil
}

func (repo *TripRepo) GetAllTrips() ([]models.Trip, error) {
	var trips []models.Trip
	repo.db.Find(&trips)
	return trips, nil
}

func (repo *TripRepo) GetTripById(id uint) (models.Trip, error) {
	var trip models.Trip
	repo.db.Find(&trip, "trips_id = ?", id)
	return trip, nil
}
