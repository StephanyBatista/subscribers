package campaigns

import (
	"subscribers/domain"
)

type Campaign struct {
	*domain.Entity
	Name        string           `gorm:"size:100; not null"`
	Description string           `gorm:"size:250"`
	Active      bool             `gorm:"not null"`
	CreatedBy   domain.UserValue `gorm:"embedded;embeddedPrefix:createdby_"`
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
