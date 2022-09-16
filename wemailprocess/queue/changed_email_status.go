package queue

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"wemailprocess/data"
	"wemailprocess/queue/types"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func ListenChangedEmailStatus(queue IQueuebase, db *sql.DB, wg sync.WaitGroup) {
	queueURL := os.Getenv("AWS_URL_QUEUE_CHANGED_EMAIL_STATUS")

	for {
		msgResult, err := queue.GetMessages(queueURL)
		if err != nil {
			fmt.Println("Error svc.ReceiveMessage: ", err.Error())
		}

		if msgResult != nil {
			for _, item := range msgResult.Messages {
				updated, _ := processChangedEmailStatusMessage(item, db)
				if updated {
					queue.DeleteMessage(queueURL, *item.ReceiptHandle)
				}
			}
		}
	}
	wg.Done()
}

func processChangedEmailStatusMessage(message *sqs.Message, db *sql.DB) (bool, error) {

	var changedResponse types.ChangedEmailEventResponse
	body := strings.Replace(*message.Body, "\n", "", -1)
	err := json.Unmarshal([]byte(body), &changedResponse)
	if err != nil {
		return false, errors.New("processMessage() error json type ChangedEmailEventResponse: " + err.Error())
	}

	var messageResponse types.MessageEventResponse
	err = json.Unmarshal([]byte(changedResponse.Message), &messageResponse)
	if err != nil {
		return false, errors.New("processMessage() error json type MessageEventResponse: " + err.Error())
	}
	fmt.Println("EventType ", messageResponse.EventType)
	fmt.Println("MessageId", messageResponse.Mail.MessageId)
	updated, _ := data.UpdateStatusSubscriber(db, messageResponse.EventType, messageResponse.Mail.MessageId)
	return updated, nil
}
