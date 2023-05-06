package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateDocument(db *gorm.DB, newDocument *models.Document) error {
	db.Create(&newDocument)
	return nil
}
