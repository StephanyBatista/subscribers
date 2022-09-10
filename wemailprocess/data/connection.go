package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	log.Println("creating new connection to database")
	db, err := sql.Open("postgres", "host=localhost user=user password=@postgres dbname=subscriber_dev port=5432 sslmode=disable")

	if err != nil {
		log.Fatal("error open connection to database")
	}

	return db
}
