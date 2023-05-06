package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("abobus.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Document{}, &models.Trip{})
	if err != nil {
		panic("failed to AutoMigrate")
	}
	return db

}
