package queue

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"wemailprocess/data"
	"wemailprocess/mail"
	"wemailprocess/queue/types"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/rs/xid"
)

func ListenCampaignReady(queue IQueuebase, db *sql.DB, wg sync.WaitGroup) {
	queueURL := os.Getenv("AWS_URL_QUEUE_CAMPAIGN_READY")

	for {
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
	wg.Done()
}

func processCampaignReady(message *sqs.Message, session *session.Session, db *sql.DB) error {
	var response types.CampaignReadyResponse
	err := json.Unmarshal([]byte(*message.Body), &response)
	if err != nil {
		return errors.New("processCampaignReady() error json type CampaignReadyResponse: " + err.Error())
	}

	fmt.Println("processCampaignReady(): CampaignID: " + response.Id)
	campaign := data.GetCampaignBy(db, response.Id)
	if campaign.Id == "" {
		return errors.New("campaign not found")
	}
	contacts := data.GetContactsBy(db, campaign.CreatedById)

	for _, contact := range contacts {
		subscriber := data.Subscriber{
			Id:         xid.New().String(),
			Email:      contact.Email,
			CampaignID: response.Id,
			ContactID:  contact.Id,
			Status:     "Waiting",
		}
		providerKey := mail.Send(session, subscriber.Email, campaign.Subject, campaign.Body)
		subscriber.ProviderEmailKey = providerKey
		data.SaveSubscriber(db, subscriber)
		fmt.Println("processCampaignReady(): Save contact as subscriber: " + contact.Name)
	}

	return nil
}
