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

func (repo *DocumentRepo) GetDocumentById(id uint) (models.Document, error) {
	var document models.Document
	repo.db.Find(&document, "document_id = ?", id)
	return document, nil
}

func (repo *DocumentRepo) UpdateDocument(document *models.Document, name string, description string, discount float64) (models.Document, error) {
	repo.db.Model(&document).UpdateColumns(models.Document{Name: name, Description: description, Discount: discount})
	return *document, nil
}
