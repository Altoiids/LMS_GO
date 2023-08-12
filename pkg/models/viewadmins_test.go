package models

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestViewAdmins(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "email"}).
		AddRow("admin", "admin@gmail.com").
		AddRow("altoids", "al@gmail.com")

	mock.ExpectQuery("SELECT name, email FROM user where adminId = 1").
		WillReturnRows(rows)

	admins, error := ViewAdmins(db)

	assert.Equal(t, nil, error)
	assert.Equal(t, 2, len(admins))

	assert.Equal(t, "admin", admins[0].Name)
	assert.Equal(t, "admin@gmail.com", admins[0].Email)
	assert.Equal(t, "altoids", admins[1].Name)
	assert.Equal(t, "al@gmail.com", admins[1].Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
