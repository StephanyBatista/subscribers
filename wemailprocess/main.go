package main

import (
	"sync"
	"wemailprocess/data"
	"wemailprocess/queue"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := data.GetDB()
	session := queue.NewSession()
	queueBase := queue.QueueBase{Session: session}

	var wg sync.WaitGroup
	wg.Add(2)

	go queue.ListenCampaignReady(&queueBase, db, wg)
	go queue.ListenChangedEmailStatus(&queueBase, db, wg)

	wg.Wait()
}
