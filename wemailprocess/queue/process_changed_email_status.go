package queue

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"wemailprocess/data"
	"wemailprocess/queue/types"
)

func processChangedEmailStatusMessage(message string, db *sql.DB) (bool, error) {

	var changedResponse types.ChangedEmailEventResponse
	body := strings.Replace(message, "\n", "", -1)
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
