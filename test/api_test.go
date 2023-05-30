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
	var tripsService = services.NewTripsService(tripsRepo)

	r := gin.Default()
	r.GET("/trips", router2.GetAllTrips(tripsService))
	r.POST("/trips", router2.CreateTrip(tripsService))

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New().
		Handler(r).
		Get("/trips").
		//JSON(`{"email": "a@b.c", "password": "11111111"}`).
		Expect(t).
		//CookiePresent("token").
		Status(http.StatusOK).
		End()
}
