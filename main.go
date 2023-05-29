package main

import (
	"github.com/gin-gonic/gin"
	database "ukrabobus/db"
	router2 "ukrabobus/handlers/router"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func main() {
	r := gin.Default()

	db := database.CreateDB()
	var authService = services.NewAuthService(db)
	var tripsRepo = repos.NewTripRepo(db)
	var docsRepo = repos.NewDocumentRepo(db)
	var ticketsRepo = repos.NewTicketRepo(db)
	var usersRepo = repos.NewUserRepo(db)

	var docsService = services.NewDocsService(docsRepo)
	var tripsService = services.NewTripsService(tripsRepo)
	var ticketsService = services.NewTicketsService(ticketsRepo)

	r.GET("/trips", router2.GetAllTrips(tripsService))
	r.POST("/trips", router2.CreateTrip(tripsService))
	r.GET("/users", router2.GetAllUsers(usersRepo))
	r.POST("/users", router2.CreateUser(usersRepo))
	r.GET("/documents", router2.GetAllDocuments(docsService))
	r.POST("/documents", router2.CreateDocument(docsService))
	r.GET("/tickets", router2.GetAllTickets(ticketsService))
	r.POST("/tickets", router2.CreateTicket(ticketsService))
	r.POST("/login", router2.Login(authService))
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
