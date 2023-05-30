package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ukrabobus/models"
)

func TestCreateTokenByService(t *testing.T) {
	db, _ := createMockDB(t)
	user := &models.User{
		FirstName:  "1",
		LastName:   "2",
		Email:      "3",
		Password:   "4",
		IsAdmin:    false,
		DocumentID: 1,
	}
	_, _, err := NewAuthService(db).CreateToken(user)
	assert.NoError(t, err)
}
