package test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"testing"
	"time"
	"ukrabobus/models"
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

func TestTripCreate(t *testing.T) {
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

	t.Run("create and get new trip", func(t *testing.T) {
		var err = tripRepo.CreateTrip(newTrip)

		if err != nil {
			t.Errorf("failed to creare trip with error: %v", err)
		}

		trip, err := tripRepo.GetTripById(newTrip.TripsID)
		if err != nil {
			t.Errorf("failed to get trip with error: %v", err)
		}

		if !reflect.DeepEqual(newTrip, &trip) {
			t.Errorf("trip data is corrupted; actual: %v, expected: %v", trip, newTrip)

		}
	})
}

func TestUserCreate(t *testing.T) {
	var db, err = CreateFakeDB(t)
	if err != nil {
		t.Errorf("failed to create fake db: %v", err)
	}

	newUser := &models.User{
		UserID:     78,
		FirstName:  "1",
		LastName:   "2",
		Email:      "3",
		Password:   "4",
		IsAdmin:    false,
		DocumentID: 1,
	}

	var userRepo = repos.NewUserRepo(db)

	t.Run("create and get new user", func(t *testing.T) {
		var err = userRepo.CreateUser(newUser)

		if err != nil {
			t.Errorf("failed to creare user with error: %v", err)
		}

		err = userRepo.DeleteUser(newUser)
		if err != nil {
			t.Errorf("failed to delete user with error: %v", err)
		}
		var user2 models.User

		user2, err = userRepo.GetUserById(newUser.UserID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
		}
		if user2.UserID != 0 {
			t.Errorf("failed to delete user with error: %v", err)
		}

	})
}

func TestDocumentUpdate(t *testing.T) {
	var db, err = CreateFakeDB(t)
	if err != nil {
		t.Errorf("failed to create fake db: %v", err)
	}

	newDocument := &models.Document{
		DocumentID:  26,
		Name:        "testName",
		Description: "testDescription",
		Discount:    0.1,
	}

	var documentRepo = repos.NewDocumentRepo(db)

	t.Run("create and get new document", func(t *testing.T) {
		var err = documentRepo.CreateDocument(newDocument)

		if err != nil {
			t.Errorf("failed to creare document with error: %v", err)
		}

		document, err := documentRepo.UpdateDocument(newDocument, "dd", "ff", 1)
		if err != nil {
			t.Errorf("failed to update document with error: %v", err)
		}

		if reflect.DeepEqual(newDocument, document) {
			t.Errorf("document data is corrupted; actual: %v, expected: %v", document, newDocument)

		}
	})
}
