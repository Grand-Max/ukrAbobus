package service

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

func TestCreateUserByService(t *testing.T) {
	db, mock := createMockDB(t)
	user := &models.User{
		FirstName:  "1",
		LastName:   "2",
		Email:      "3",
		Password:   "4",
		IsAdmin:    false,
		DocumentID: 1,
	}
	mock.ExpectBegin()
	const sqlInsert = "INSERT INTO `users` (`first_name`,`last_name`,`email`,`password`,`is_admin`, `document_id`) VALUES (?,?,?,?,?,?)"
	var s = regexp.QuoteMeta(sqlInsert)
	fmt.Println(s)
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs("1", "2", "3", "4", "false", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := NewUserService(repos.NewUserRepo(db)).CreateUser(*user)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsUserOkByService(t *testing.T) {
	db, _ := createMockDB(t)
	userOK := &models.User{
		FirstName:  "1",
		LastName:   "2",
		Email:      "3",
		Password:   "4",
		IsAdmin:    false,
		DocumentID: 1,
	}
	isOk := NewUserService(repos.NewUserRepo(db)).IsUserOk(*userOK)
	assert.Equal(t, true, isOk)
}
