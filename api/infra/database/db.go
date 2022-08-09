package database

import (
	"os"
	"subscribers/domain/campaigns"
	"subscribers/domain/clients"
	"subscribers/domain/users"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	db := getDb()
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&clients.Client{})
	db.AutoMigrate(&campaigns.Campaign{})
	db.AutoMigrate(&campaigns.Subscriber{})

	return db
}

func getDb() *gorm.DB {
	connectionString := os.Getenv("sub_database")

	var db *gorm.DB
	var err error
	if connectionString == "sqlite:memory" {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	} else if connectionString == "sqlite" || connectionString == "" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
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
