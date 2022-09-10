package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
)

//get access key on https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html

const (
	CharSet = "UTF-8"
)

func Send(session *session.Session, recipient, subject, textBody string) string {

	svc := ses.New(session)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String("This is an automatic email"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(os.Getenv("AWS_EMAIL_SENDER")),
		// Uncomment to use a configuration set
		ConfigurationSetName: aws.String(os.Getenv("AWS_CONFIGURATION_SET")),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	return *result.MessageId
}
