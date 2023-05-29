package service

import (
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

type TripService struct {
	repo *repos.TripRepo
}

func NewTripsService(tripRepo *repos.TripRepo) *TripService {
	return &TripService{
		repo: tripRepo,
	}
}

func (service *TripService) CreateTrip(newTrip models.Trip) error {

	err := service.repo.CreateTrip(&newTrip)
	return err

}

func (service *TripService) IsTripOk(trip models.Trip) bool {
	var isOk = true

	if trip.CityTo == "" ||
		trip.CityFrom == "" ||
		trip.ArrivalTime.IsZero() ||
		trip.DepartureTime.IsZero() ||
		trip.NumberOfPlaces == 0 ||
		trip.TransportType == "" {
		isOk = false
	}

	return isOk
}
