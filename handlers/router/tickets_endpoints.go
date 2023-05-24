package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllTickets(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ticketRepo = repos.NewTicketRepo(db)
		tickets, err := ticketRepo.GetAllTickets()
		if err != nil {
			ctx.JSON(200, tickets)
			return
		}
		ctx.Status(500)
	}
}

func CreateTicket(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTicket models.Ticket
		var ticketRepo = repos.NewTicketRepo(db)

		if err := ctx.BindJSON(&newTicket); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		if !services.IsTicketOk(newTicket) {
			ctx.Status(400)
			return
		}
		err := ticketRepo.CreateTicket(&newTicket)
		if err != nil {
			ctx.Status(500)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, newTicket)
	}
}
