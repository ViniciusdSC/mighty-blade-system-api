package gorm_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateGormMock(t *testing.T) *gorm.DB {
	mDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mDB,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db
}
