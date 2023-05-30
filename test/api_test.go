package test

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/steinfletcher/apitest"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	router2 "ukrabobus/handlers/router"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func CreateFakeDB2() (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	err = gormDB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Document{}, &models.Trip{})

	return gormDB, err
}

func TestAll(t *testing.T) {
	db, err := CreateFakeDB2()
	if err != nil {
		t.Errorf("failed to create fake db: %v", err)
	}

	var tripsRepo = repos.NewTripRepo(db)
	var docsRepo = repos.NewDocumentRepo(db)
	var usersRepo = repos.NewUserRepo(db)

	var docsService = services.NewDocsService(docsRepo)
	var tripsService = services.NewTripsService(tripsRepo)
	var userService = services.NewUserService(usersRepo)

	r := gin.Default()
	r.GET("/trips", router2.GetAllTrips(tripsService))
	r.POST("/trips", router2.CreateTrip(tripsService))
	r.GET("/users", router2.GetAllUsers(userService))
	r.POST("/users", router2.CreateUser(userService))
	r.GET("/documents", router2.GetAllDocuments(docsService))
	r.POST("/documents", router2.CreateDocument(docsService))

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Handler(r).
		Post("/trips").
		JSON("{ \"CityFrom\": \"Horokhiv\", \"CityTo\": \"Lviv\", \"ArrivalTime\": \"2023-05-06T00:00:00Z\", \"DepartureTime\": \"2023-05-07T00:00:00Z\", \"NumberOfPlaces\": 50, \"TransportType\": \"Train\"}").
		Expect(t).
		Status(http.StatusCreated).
		End()

	apitest.New().
		Handler(r).
		Get("/users").
		Expect(t).
		Status(http.StatusOK).
		End()
	apitest.New().
		Handler(r).
		Get("/documents").
		Expect(t).
		Status(http.StatusOK).
		End()
}
