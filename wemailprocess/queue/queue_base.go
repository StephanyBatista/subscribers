package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type IQueuebase interface {
	GetSession() *session.Session
	GetMessages(queueURL string) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(queueUrl string, receiptHandle string)
}

type QueueBase struct {
	Session *session.Session
	SQS     *sqs.SQS
}

func (q *QueueBase) GetSession() *session.Session {
	return q.Session
}

func (q *QueueBase) GetMessages(queueURL string) (*sqs.ReceiveMessageOutput, error) {
	if q.SQS == nil {
		q.SQS = sqs.New(q.Session)
	}

	msgResult, err := q.SQS.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	})
	return msgResult, err
}

func (q *QueueBase) DeleteMessage(queueURL string, receiptHandle string) {

	q.SQS.DeleteMessage(
		&sqs.DeleteMessageInput{QueueUrl: aws.String(queueURL), ReceiptHandle: aws.String(receiptHandle)})
}
