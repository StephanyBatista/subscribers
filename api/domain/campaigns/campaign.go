package campaigns

import (
	"subscribers/domain"
)

const (
	Waiting string = "Waiting"
	Draft          = "Draft"
	Sending        = "Processing"
	Sent           = "Sent"
	NotSent        = "NotSent"
	Read           = "Read"
)

type Campaign struct {
	domain.Entity
	Name      string           `json:"name" gorm:"size:100; not null"`
	Status    string           `json:"status" gorm:"size:15; not null"`
	From      string           `json:"from" gorm:"size:100; not null"`
	Subject   string           `json:"subject" gorm:"size:150;not null"`
	Body      string           `json:"body" gorm:"not null"`
	CreatedBy domain.UserValue `json:"createdBy" gorm:"embedded;embeddedPrefix:createdby_"`
}

func (c *Campaign) Sending() {
	c.Status = Sending
}

func NewCampaign(name, from, subject, body, userId, userName string) *Campaign {

	return &Campaign{
		Name:      name,
		Status:    Draft,
		From:      from,
		Subject:   subject,
		Body:      body,
		Entity:    domain.NewEntity(),
		CreatedBy: domain.UserValue{Id: userId, Name: userName},
	}
}
