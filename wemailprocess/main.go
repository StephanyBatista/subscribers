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
	queueBase := queue.QueueBase{Session: session}

	go queue.ListenCampaignReady(&queueBase, db)
	go queue.ListenChangedEmailStatus(&queueBase, db)

	for {
	}
}
