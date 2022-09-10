package queue

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func Test_processMessage_must_return_error_when_message_body_has_error(t *testing.T) {

	output := &sqs.Message{Body: aws.String("invalid json")}

	_, err := processChangedEmailStatusMessage(output, nil)

	assert.Contains(t, err.Error(), "processMessage() error json type ChangedEmailEventResponse")
}

func Test_processMessage_must_return_error_when_message_event_response_has_error(t *testing.T) {

	json := `{ "Type": "Notification", "Message": "invalid json"}`
	output := &sqs.Message{Body: aws.String(json)}

	_, err := processChangedEmailStatusMessage(output, nil)

	assert.Contains(t, err.Error(), "processMessage() error json type MessageEventResponse")
}

func Test_processMessage_must_update_status_subscriber(t *testing.T) {

	json := `{ "Type": "Notification", "Message": "{\"eventType\": \"Open\", \"mail\": {\"messageId\": \"xpti\"}}"}`
	output := &sqs.Message{Body: aws.String(json)}
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("UPDATE subscribers").ExpectExec().WithArgs("Open", "xpti").WillReturnResult(sqlmock.NewResult(1, 1))

	updated, _ := processChangedEmailStatusMessage(output, db)

	assert.True(t, updated)
}

func Test_processMessage_must_not_update_status_subscriber_when_not_found_by_message_id(t *testing.T) {

	json := `{ "Type": "Notification", "Message": "{\"eventType\": \"Open\", \"mail\": {\"messageId\": \"xpti\"}}"}`
	output := &sqs.Message{Body: aws.String(json)}
	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("UPDATE subscribers").ExpectExec().WithArgs("Open", "xpti").WillReturnResult(sqlmock.NewResult(0, 0))

	updated, _ := processChangedEmailStatusMessage(output, db)

	assert.False(t, updated)
}
