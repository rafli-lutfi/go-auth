package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate The Schema
	db.AutoMigrate()
}

func GetDBConnection() *gorm.DB {
	return db
}
