package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"ukrabobus/models"
	"ukrabobus/router"
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
	r.GET("/users", router.GetAllUsers(db))
	r.POST("/users", router.CreateUser(db))
	r.GET("/documents", router.GetAllDocuments(db))
	r.POST("/documents", router.CreateDocument(db))
	r.GET("/tickets", router.GetAllTickets(db))
	r.POST("/tickets", router.CreateTicket(db))
	r.Run() // listen and serve on 0.0.0.0:8080
}
