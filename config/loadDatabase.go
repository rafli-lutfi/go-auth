package config

import (
	"os"

	"github.com/rafli-lutfi/go-auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate The Schema
	db.AutoMigrate(model.User{})
}

func GetDBConnection() *gorm.DB {
	return db
}
