package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestRepository_Save(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	imageMockRows := sqlmock.NewRows([]string{"id"}).
		AddRow("2")

	mock.ExpectQuery("INSERT INTO orders").
		WithArgs("OrderId_sample", "defId_sample").
		WillReturnRows(imageMockRows)

	a := New(sqlxDB)

	if _, err := a.Save("OrderId_sample", "defId_sample"); err != nil {
		t.Errorf("error was not expected while saving data: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
