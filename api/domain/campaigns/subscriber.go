package campaigns

import "subscribers/domain"

type Subscriber struct {
	domain.Entity
	Campaign         Campaign
	CampaignID       string `json:"campaignID" gorm:"size:25;not null"`
	ContactID        string `json:"contactID" gorm:"size:25;not null"`
	Email            string `json:"email" gorm:"size:100;not null"`
	Status           string `json:"status" gorm:"size:15;not null"`
	ProviderEmailKey string `json:providerEmailKey gorm:"size:25"`
}

func (s *Subscriber) Sent() {
	s.Status = Sent
}

func (s *Subscriber) NotSent() {
	s.Status = NotSent
}

func (s *Subscriber) Read() {
	s.Status = Read
}

func NewSubscriber(campaign Campaign, clientID string, email string) *Subscriber {
	return &Subscriber{
		Entity:    domain.NewEntity(),
		Campaign:  campaign,
		ContactID: clientID,
		Email:     email,
		Status:    Waiting,
	}
}
