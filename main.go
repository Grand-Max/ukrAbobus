package main

import (
	"github.com/gin-gonic/gin"
	database "ukrabobus/db"
	router2 "ukrabobus/handlers/router"
)

func main() {
	r := gin.Default()

	db := database.CreateDB()

	r.GET("/trips", router2.GetAllTrips(db))
	r.POST("/trips", router2.CreateTrip(db))
	r.GET("/users", router2.GetAllUsers(db))
	r.POST("/users", router2.CreateUser(db))
	r.GET("/documents", router2.GetAllDocuments(db))
	r.POST("/documents", router2.CreateDocument(db))
	r.GET("/tickets", router2.GetAllTickets(db))
	r.POST("/tickets", router2.CreateTicket(db))
	r.POST("/login", router2.Login(db))
	r.Run() // listen and serve on 0.0.0.0:8080
}
