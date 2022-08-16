package mocks

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDatabase() (*gorm.DB, sqlmock.Sqlmock, error) {

	// get db and mock
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}

	// create dialector
	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	// open the database
	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return db, mock, err
}
