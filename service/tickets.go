package service

import "ukrabobus/models"

func IsTicketOk(trip models.Ticket) bool {
	var isOk = true

	if trip.UserID == 0 ||
		trip.TripID == 0 ||
		trip.Price == 0 {
		isOk = false
	}

	return isOk
}
