package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	migrate "github.com/rubenv/sql-migrate"
)

func ApplyMigration(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "commun/database/migrations",
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
