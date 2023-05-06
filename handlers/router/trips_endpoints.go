package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ukrabobus/models"
	tripsService "ukrabobus/service"
)

func GetAllTrips(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var trips []models.Trip
		db.Find(&trips)
		ctx.JSON(200, trips)
	}
}

func CreateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTrip models.Trip

		if err := ctx.BindJSON(&newTrip); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		if tripsService.IsTripOk(newTrip) {
			db.Create(&newTrip)
			ctx.IndentedJSON(http.StatusCreated, newTrip)
		} else {
			ctx.Status(400)
		}

	}
}