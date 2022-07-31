package database

import (
	"os"
	"subscribers/domain/campaigns"
	"subscribers/domain/users"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	db := getDb()
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&campaigns.Campaign{})

	return db
}

func getDb() *gorm.DB {
	connectionString := os.Getenv("sub_database")
	if connectionString == "" {
		panic("enviroment sub_database is not filled")
	}
	var db *gorm.DB
	var err error
	if connectionString == "sqlite" {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  connectionString,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	}

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
