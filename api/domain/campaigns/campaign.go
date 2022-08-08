package campaigns

import (
	"subscribers/domain"
)

type Campaign struct {
	domain.Entity
	Name      string           `json:"name" gorm:"size:100; not null"`
	From      string           `json:"from" gorm:"size:100; not null"`
	Body      string           `json:"body" gorm:"not null"`
	CreatedBy domain.UserValue `json:"createdBy" gorm:"embedded;embeddedPrefix:createdby_"`
}

func NewCampaign(name, from, body, userId, userName string) *Campaign {

	return &Campaign{
		Name:      name,
		From:      from,
		Body:      body,
		Entity:    domain.NewEntity(),
		CreatedBy: domain.UserValue{Id: userId, Name: userName},
	}
}
