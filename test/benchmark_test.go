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

func CreateBenchmarkFakeDB() (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	err = gormDB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Document{}, &models.Trip{})

	return gormDB, err
}

func BenchmarkInsert(b *testing.B) {
	db, err := CreateBenchmarkFakeDB()
	if err != nil {
		b.Errorf("failed to create fake db: %v", err)
	}

	var tripsRepo = repos.NewTripRepo(db)
	var tripsService = services.NewTripsService(tripsRepo)

	r := gin.Default()
	r.GET("/trips", router2.GetAllTrips(tripsService))
	r.POST("/trips", router2.CreateTrip(tripsService))

	ts := httptest.NewServer(r)
	defer ts.Close()

	for i := 0; i < b.N; i++ {
		apitest.New().
			Handler(r).
			Post("/trips").
			JSON("{ \"CityFrom\": \"Horokhiv\", \"CityTo\": \"Lviv\", \"ArrivalTime\": \"2023-05-06T00:00:00Z\", \"DepartureTime\": \"2023-05-07T00:00:00Z\", \"NumberOfPlaces\": 50, \"TransportType\": \"Train\"}").
			Expect(b).
			Status(http.StatusCreated).
			End()
	}
}
