package service

import (
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

type TicketService struct {
	repo *repos.TicketRepo
}

func NewTicketsService(ticketRepo *repos.TicketRepo) *TicketService {
	return &TicketService{
		repo: ticketRepo,
	}
}

func (service *TicketService) CreateTicket(newTicket models.Ticket) error {

	err := service.repo.CreateTicket(&newTicket)
	return err

}

func (service *TicketService) GetAllTickets() ([]models.Ticket, error) {
	tickets, err := service.repo.GetAllTickets()
	return tickets, err
}

func (service *TicketService) IsTicketOk(ticket models.Ticket) bool {
	var isOk = true

	if ticket.UserID == 0 ||
		ticket.TripID == 0 ||
		ticket.Price == 0 {
		isOk = false
	}

	return isOk
}
