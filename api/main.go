package main

import (
	"subscribers/infra/database"
	"subscribers/modules/contacts"
	"subscribers/modules/users"
	"subscribers/modules/web"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	loadEnvs()
	db := database.GetConnection()
	database.ApplyMigration(db)
	r := web.CreateRouter()
	users.ApplyRouter(r, db)
	contacts.ApplyRouter(r, db)

	r.Run(":6004")
}

func loadEnvs() {
	godotenv.Load()
}
