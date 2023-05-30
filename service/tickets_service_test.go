package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

func TestCreateTicketByService(t *testing.T) {
	db, mock := createMockDB(t)
	ticket := &models.Ticket{
		UserID: 1,
		TripID: 1,
		Price:  222.5,
	}
	mock.ExpectBegin()
	const sqlInsert = "INSERT INTO `tickets` (`user_id`,`trip_id`,`price`) VALUES (?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs(1, 1, 222.5).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := NewTicketsService(repos.NewTicketRepo(db)).CreateTicket(*ticket)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsTicketOkByService(t *testing.T) {
	db, _ := createMockDB(t)
	ticketOK := &models.Ticket{
		UserID: 1,
		TripID: 1,
		Price:  222.5,
	}
	isOk := NewTicketsService(repos.NewTicketRepo(db)).IsTicketOk(*ticketOK)
	assert.Equal(t, true, isOk)
}
