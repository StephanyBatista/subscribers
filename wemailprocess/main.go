package main

import (
	"wemailprocess/data"
	"wemailprocess/queue"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := data.GetDB()
	session := queue.NewSession()
	queueBase := queue.Queue{Session: session}

	queue.Listen(&queueBase, db)
}
