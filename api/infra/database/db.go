package database

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
)

func ApplyMigration(db *sql.DB) {
	//db := getDb()
	//db.AutoMigrate(&users.User{})

	migrations := &migrate.FileMigrationSource{
		Dir: "infra/database/migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("ApplyMigration(): migrate.Exec ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("sub_database"))
	if err != nil {
		log.Fatal("ApplyMigration(): sql.Open ", err)
	}
	return db
}
