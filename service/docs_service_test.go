package service

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
	"ukrabobus/models"
	repos "ukrabobus/repository"
)

func createMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	var db *sql.DB
	var err error

	db, mock, err := sqlmock.New() // mock sql.DB
	assert.NoError(t, err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	assert.NoError(t, err)

	return gormDB, mock
}

func TestCreateDocByService(t *testing.T) {
	db, mock := createMockDB(t)
	doc := &models.Document{
		Name:        "testName",
		Description: "testDescription",
		Discount:    0.1,
	}
	mock.ExpectBegin()
	const sqlInsert = "INSERT INTO `documents` (`name`,`description`,`discount`) VALUES (?,?,?)"
	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).
		WithArgs("testName", "testDescription", 0.1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := NewDocsService(repos.NewDocumentRepo(db)).CreateDocument(*doc)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsDocOkByService(t *testing.T) {
	db, _ := createMockDB(t)
	docOK := &models.Document{
		Name:        "testName",
		Description: "testDescription",
		Discount:    0.1,
	}
	isOk := NewDocsService(repos.NewDocumentRepo(db)).IsDocumentOk(*docOK)
	assert.Equal(t, true, isOk)
}
