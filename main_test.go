package main

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// a successful case
func TestShouldGetNameByAge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Tim").
		AddRow("Tam")

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE age = (.+)$").WithArgs(1).WillReturnRows(rows)

	// now we execute our method
	if _, err = getNameByAge(1, db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
