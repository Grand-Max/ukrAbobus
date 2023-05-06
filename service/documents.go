package service

import "ukrabobus/models"

func IsDocumentOk(document models.document) bool {
	var isOk = true

	if document.Name == "" ||
		document.Description == "" ||
		document.Discount == 0 {
		isOk = false
	}

	return isOk
}