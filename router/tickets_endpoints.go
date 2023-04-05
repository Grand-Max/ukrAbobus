package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
)

func GetAllTickets(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tickets []models.Ticket
		db.Find(&tickets)
		ctx.JSON(200, tickets)
	}
}

func CreateTicket(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTicket models.Ticket

		if err := ctx.BindJSON(&newTicket); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		db.Create(&newTicket)
		ctx.IndentedJSON(http.StatusCreated, newTicket)
	}
}
