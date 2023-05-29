package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ukrabobus/models"
	services "ukrabobus/service"
)

func GetAllTickets(ticketRepo *services.TicketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tickets, err := ticketRepo.GetAllTickets()
		if err != nil {
			ctx.JSON(200, tickets)
			return
		}
		ctx.Status(500)
	}
}

func CreateTicket(service *services.TicketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTicket models.Ticket

		if err := ctx.BindJSON(&newTicket); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}
		if !service.IsTicketOk(newTicket) {
			ctx.Status(400)
			return
		}
		err := service.CreateTicket(newTicket)
		if err != nil {
			ctx.Status(500)
			return
		}
		ctx.IndentedJSON(http.StatusCreated, newTicket)
	}
}
