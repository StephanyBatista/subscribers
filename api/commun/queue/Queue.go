package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type IQueue interface {
	Send(queueURL, body string) error
}

type Queue struct{}

func (q *Queue) Send(queueURL, body string) error {
	session := newSession()
	sqsClient := sqs.New(session)
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(body),
	})
	return err
}
