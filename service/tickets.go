package service

import "ukrabobus/models"

func IsTripOk(trip models.Trip) bool {
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
