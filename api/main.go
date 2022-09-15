package main

import (
	"subscribers/modules/campaigns"
	"subscribers/modules/contacts"
	"subscribers/modules/files"
	"subscribers/modules/healtchcheck"
	"subscribers/modules/users"
	"subscribers/utils/database"
	"subscribers/utils/queue"
	"subscribers/utils/webtest"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	loadEnvs()
	db := database.GetConnection()
	session := queue.NewSession()
	database.ApplyMigration(db)
	r := webtest.CreateRouter()
	users.ApplyRouter(r, db)
	contacts.ApplyRouter(r, db)
	campaigns.ApplyRouter(r, db, session)
	healtchcheck.ApplyRouter(r, db)
	files.ApplyRouter(r, db)

	r.Run(":6004")
}

func loadEnvs() {
	godotenv.Load()
}
