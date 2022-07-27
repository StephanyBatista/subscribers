package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	connectionString := os.Getenv("sub_database")
	if connectionString == "" {
		panic("enviroment sub_database is not filled")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
