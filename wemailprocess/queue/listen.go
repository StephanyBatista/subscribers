package queue

import (
	"database/sql"
	"fmt"
	"os"
)

func Listen(queue IQueue, db *sql.DB) {

	for {
		listenChangeEmailStatus(queue, db)
		listenCampaignReady(queue, db)
	}
}

func listenChangeEmailStatus(queue IQueue, db *sql.DB) {
	queueURL := os.Getenv("AWS_URL_QUEUE_CHANGED_EMAIL_STATUS")
	msgResult, err := queue.GetMessages(queueURL)
	if err != nil {
		fmt.Println("Error svc.ReceiveMessage: ", err.Error())
	}

	if msgResult != nil {
		for _, item := range msgResult.Messages {
			updated, _ := processChangedEmailStatusMessage(*item.Body, db)
			if updated {
				queue.DeleteMessage(queueURL, *item.ReceiptHandle)
			}
		}
	}
}

func listenCampaignReady(queue IQueue, db *sql.DB) {
	queueURL := os.Getenv("AWS_URL_QUEUE_CAMPAIGN_READY")

	msgResult, err := queue.GetMessages(queueURL)
	if err != nil {
		fmt.Println("Error svc.ReceiveMessage: ", err.Error())
	}

	if msgResult != nil {
		for _, item := range msgResult.Messages {
			fmt.Println("ListenCampaignReady(): A new message arrived")
			err := processCampaignReady(item, queue.GetSession(), db)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				queue.DeleteMessage(queueURL, *item.ReceiptHandle)
			}
		}
	}
}
