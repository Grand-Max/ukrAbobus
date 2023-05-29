package service

import (
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

type DocumentService struct {
	repo *repos.DocumentRepo
}

func NewDocsService(docRepo *repos.DocumentRepo) *DocumentService {
	return &DocumentService{
		repo: docRepo,
	}
}

func (service *DocumentService) CreateDocument(newDocument models.Document) error {

	err := service.repo.CreateDocument(&newDocument)
	return err

}

func (service *DocumentService) IsDocumentOk(document models.Document) bool {
	var isOk = true

	if document.Name == "" ||
		document.Description == "" ||
		document.Discount == 0 {
		isOk = false
	}

	return isOk
}
