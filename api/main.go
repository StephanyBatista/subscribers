package main

import (
	"subscribers/commun/database"
	"subscribers/commun/queue"
	"subscribers/commun/webtest"
	"subscribers/modules/campaigns"
	"subscribers/modules/contacts"
	"subscribers/modules/files"
	"subscribers/modules/healtchcheck"
	"subscribers/modules/users"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	loadEnvs()
	db := database.GetConnection()
	database.ApplyMigration(db)
	r := webtest.CreateRouter()
	users.ApplyRouter(r, db)
	contacts.ApplyRouter(r, db)
	campaigns.ApplyRouter(r, db, &queue.Queue{})
	healtchcheck.ApplyRouter(r, db)
	files.ApplyRouter(r, db)

	r.Run(":6004")
}

func loadEnvs() {
	godotenv.Load()
}
