package clients

import "subscribers/domain"

type Client struct {
	domain.Entity
	Name      string           `json:"name" gorm:"not null;size:100;"`
	Email     string           `json:"email" gorm:"not null;size:100;"`
	Active    bool             `json:"active" gorm:"not null"`
	CreatedBy domain.UserValue `json:"createdBy" gorm:"embedded;embeddedPrefix:createdby_"`
}

func NewClient(request CreationRequest, user *domain.UserValue) (*Client, []error) {
	errs := domain.Validate(request)
	if errs != nil {
		return nil, errs
	}
	client := &Client{
		Name:   request.Name,
		Email:  request.Email,
		Active: true,
		Entity: domain.NewEntity(),
	}

	if user != nil {
		client.CreatedBy = *user
	}

	return client, nil
}
