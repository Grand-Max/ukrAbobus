package test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"testing"
	"time"
	models "ukrabobus/models"
	repos "ukrabobus/repository"
)

func createMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	var db *sql.DB
	var err error

	db, mock, err := sqlmock.New() // mock sql.DB
	assert.NoError(t, err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	return gormDB, mock
}

func CreateFakeDB(t *testing.T) (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Errorf("failed to create fake db: %v", err)
	}
	err = gormDB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Document{}, &models.Trip{})

	return gormDB, err
}

func TestGetCreate(t *testing.T) {
	var db, err = CreateFakeDB(t)
	if err != nil {
		t.Errorf("failed to create fake db: %v", err)
	}

	newTrip := &models.Trip{
		TripsID:        999,
		CityFrom:       "Kyiv",
		CityTo:         "Sevastopol",
		ArrivalTime:    time.Time{},
		DepartureTime:  time.Time{},
		NumberOfPlaces: 48,
		TransportType:  "Plane",
	}

	var tripRepo = repos.NewTripRepo(db)

	t.Run("create and get new order", func(t *testing.T) {
		var err = tripRepo.CreateTrip(newTrip)

		if err != nil {
			t.Errorf("failed to creare order with error: %v", err)
		}

		fmt.Println(tripRepo.GetAllTrips())
		fmt.Println(newTrip)
		trip, err := tripRepo.GetTripById(newTrip.TripsID)
		if err != nil {
			t.Errorf("failed to get order with error: %v", err)
		}
		fmt.Println(trip)

		if !reflect.DeepEqual(newTrip, &trip) {
			t.Errorf("order data is corrupted; actual: %v, expected: %v", trip, newTrip)

		}
	})
}
