package main

import (
	"awesomeProject1/models"
	"awesomeProject1/router"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("abobus.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Document{}, &models.Trip{})
	if err != nil {
		panic("failed to AutoMigrate")
	}

	r.GET("/trips", router.GetAllTrips(db))
	r.POST("/trips", router.CreateTrip(db))
	r.GET("/user", router.GetAllTrips(db))
	r.POST("/user", router.CreateTrip(db))
	r.GET("/trips", router.GetAllTrips(db))
	r.POST("/trips", router.CreateTrip(db))
	r.GET("/trips", router.GetAllTrips(db))
	r.POST("/trips", router.CreateTrip(db))
	r.Run() // listen and serve on 0.0.0.0:8080
}
