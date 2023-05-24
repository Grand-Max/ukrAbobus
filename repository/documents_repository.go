package repository

import (
	"gorm.io/gorm"
	"ukrabobus/models"
)

type DocumentRepo struct {
	db *gorm.DB
}

func NewDocumentRepo(database *gorm.DB) *DocumentRepo {
	return &DocumentRepo{
		db: database,
	}
}

func (repo *DocumentRepo) CreateDocument(newDocument *models.Document) error {
	repo.db.Create(&newDocument)
	return nil
}

func (repo *DocumentRepo) GetAllDocuments() ([]models.Document, error) {
	var documents []models.Document
	repo.db.Find(&documents)
	return documents, nil
}
