package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateTicket(db *gorm.DB, newTicket *models.Ticket) error {
	db.Create(&newTicket)
	return nil
}

func GetAllTickets(db *gorm.DB) ([]models.Ticket, error) {
	var tickets []models.Ticket
	db.Joins("User").Joins("Trip").Find(&tickets)
	return tickets, nil
}