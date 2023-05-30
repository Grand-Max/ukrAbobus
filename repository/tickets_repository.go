package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

type TicketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(database *gorm.DB) *TicketRepo {
	return &TicketRepo{
		db: database,
	}
}

func (repo *TicketRepo) CreateTicket(newTicket *models.Ticket) error {
	repo.db.Create(&newTicket)
	return nil
}

func (repo *TicketRepo) GetAllTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	repo.db.Joins("User").Joins("Trip").Find(&tickets)
	return tickets, nil
}

func (repo *TicketRepo) GetTicketById(id uint) (models.Ticket, error) {
	var ticket models.Ticket
	repo.db.Joins("User").Joins("Trip").Find(&ticket, "ticket_id = ?", id)
	return ticket, nil
}
