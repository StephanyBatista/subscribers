package campaigns

import (
	"subscribers/domain"
	"subscribers/modules/campaigns"
)

type Subscriber struct {
	domain.Entity
	Campaign         campaigns.Campaign
	CampaignID       string `json:"campaignID" gorm:"size:25;not null"`
	ContactID        string `json:"contactID" gorm:"size:25;not null"`
	Email            string `json:"email" gorm:"size:100;not null"`
	Status           string `json:"status" gorm:"size:15;not null"`
	ProviderEmailKey string `json:providerEmailKey gorm:"size:25"`
}

func (s *Subscriber) Sent() {
	s.Status = campaigns.Sent
}

func (s *Subscriber) NotSent() {
	s.Status = campaigns.NotSent
}

func (s *Subscriber) Read() {
	s.Status = campaigns.Read
}

func NewSubscriber(campaign campaigns.Campaign, clientID string, email string) *Subscriber {
	return &Subscriber{
		Entity:    domain.NewEntity(),
		Campaign:  campaign,
		ContactID: clientID,
		Email:     email,
		Status:    campaigns.Waiting,
	}
}
