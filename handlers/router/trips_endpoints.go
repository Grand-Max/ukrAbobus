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

func GetAllTrips(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tripsRepo = repos.NewTripRepo(db)

		trips, err := tripsRepo.GetAllTrips()
		if err != nil {
			ctx.JSON(200, trips)
			return
		}
		ctx.Status(500)
	}
}

func CreateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTrip models.Trip
		var tripsRepo = repos.NewTripRepo(db)

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
