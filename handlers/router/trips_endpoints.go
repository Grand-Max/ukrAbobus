package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ukrabobus/models"
	repos "ukrabobus/repository"
	services "ukrabobus/service"
)

func GetAllTrips(tripsRepo *repos.TripRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		trips, err := tripsRepo.GetAllTrips()
		if err != nil {
			ctx.JSON(200, trips)
			return
		}
		ctx.Status(500)
	}
}

func CreateTrip(tripsRepo *repos.TripRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTrip models.Trip

		if err := ctx.BindJSON(&newTrip); err != nil {
			fmt.Println("Bind error")
			ctx.Status(http.StatusBadRequest)
			return
		}

		if !services.IsTripOk(newTrip) {
			ctx.Status(400)
			return
		}
		err := tripsRepo.CreateTrip(&newTrip)
		if err != nil {
			ctx.Status(500)
			return
		}

	}
}
