package types

import "time"

type ChangedEmailEventResponse struct {
	Type      string
	Message   string    `json:"Message"`
	Timestamp time.Time `json:"Timestamp"`
}

type MessageEventResponse struct {
	EventType string            `json:"eventType"`
	Mail      MailEventResponse `json:"mail"`
}

type MailEventResponse struct {
	MessageId string `json:"messageId"`
	Action    string `json:"action"`
}
