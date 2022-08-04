package campaigns

import (
	"subscribers/domain"
	"subscribers/domain/campaigns/clients"
)

type Campaign struct {
	domain.Entity
	Name        string            `gorm:"size:100; not null"`
	Description string            `gorm:"size:250"`
	Active      bool              `gorm:"not null"`
	CreatedBy   domain.UserValue  `gorm:"embedded;embeddedPrefix:createdby_"`
	Clients     []*clients.Client `gorm:"many2many:campaign_clients;"`
}

func (c *Campaign) HasClient(email string) bool {
	for _, item := range c.Clients {
		if item.Email == email {
			return true
		}
	}
	return false
}

func (c *Campaign) AddClient(client *clients.Client) {
	c.Clients = append(c.Clients, client)
}

func NewCampaign(request CreationRequest, user domain.UserValue) (*Campaign, []error) {
	errs := domain.Validate(request)
	if errs != nil {
		return nil, errs
	}
	return &Campaign{
		Name:        request.Name,
		Description: request.Description,
		Active:      request.Active,
		Entity:      domain.NewEntity(),
		CreatedBy:   user,
	}, nil
}
