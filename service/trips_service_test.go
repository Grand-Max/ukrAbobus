package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

func TestCreateTripByService(t *testing.T) {
	db, mock := createMockDB(t)
	trip := &models.Trip{
		CityFrom:       "1",
		CityTo:         "2",
		ArrivalTime:    time.Time{},
		DepartureTime:  time.Time{},
		NumberOfPlaces: 48,
		TransportType:  "any",
	}
	mock.ExpectBegin()
	const sqlInsert = "INSERT INTO `trips` (`city_from`,`city_to`,`arrival_time`,`departure_time`,`number_of_places`,`transport_type`) VALUES (?,?,?,?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs("1", "2", time.Time{}, time.Time{}, 48, "any").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := NewTripsService(repos.NewTripRepo(db)).CreateTrip(*trip)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsTripOkByService(t *testing.T) {
	db, _ := createMockDB(t)
	tripOK := &models.Trip{
		CityFrom:       "1",
		CityTo:         "2",
		ArrivalTime:    time.Time{},
		DepartureTime:  time.Time{},
		NumberOfPlaces: 48,
		TransportType:  "any",
	}
	isOk := NewTripsService(repos.NewTripRepo(db)).IsTripOk(*tripOK)
	assert.Equal(t, true, isOk)
}
