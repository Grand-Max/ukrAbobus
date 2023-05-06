package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

func CreateDocument(db *gorm.DB, newDocument *models.Document) error {
	db.Create(&newDocument)
	return nil
}

func GetAllDocuments(db *gorm.DB) ([]models.Document, error) {
	var documents []models.Document
	db.Find(&documents)
	return documents, nil
}
