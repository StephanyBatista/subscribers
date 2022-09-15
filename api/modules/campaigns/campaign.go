package campaigns

import (
	"errors"
	"github.com/rs/xid"
	"net/mail"
	"time"
)

const (
	Draft          = "Draft"
	Ready          = "Ready"
	Waiting string = "Waiting"
	Sending        = "Sending"
	Sent           = "Sent"
	NotSent        = "NotSent"
)

type Campaign struct {
	Id        string    `json:"id" `
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	From      string    `json:"from"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	UserId    string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *Campaign) Ready() {
	c.Status = Ready
}

func NewCampaign(name, from, subject, body, userId string) (Campaign, []error) {

	var errs []error
	if name == "" {
		errs = append(errs, errors.New("'Name' is required"))
	}
	if from == "" {
		errs = append(errs, errors.New("'From' is required"))
	} else if _, err := mail.ParseAddress(from); err != nil {
		errs = append(errs, errors.New("'From' invalid"))
	}
	if subject == "" {
		errs = append(errs, errors.New("'Subject' is required"))
	}
	if body == "" {
		errs = append(errs, errors.New("'Body' is required"))
	}
	if userId == "" {
		errs = append(errs, errors.New("'UserId' is required"))
	}
	if len(errs) > 0 {
		return Campaign{}, errs
	}

	return Campaign{
		Id:        xid.New().String(),
		CreatedAt: time.Now().UTC(),
		Name:      name,
		Status:    Draft,
		From:      from,
		Subject:   subject,
		Body:      body,
		UserId:    userId,
	}, nil
}
