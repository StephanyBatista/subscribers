package contacts

import "subscribers/domain"

type Contact struct {
	domain.Entity
	Name   string `json:"name" gorm:"not null;size:100;"`
	Email  string `json:"email" gorm:"not null;size:100;"`
	Active bool   `json:"active" gorm:"not null"`
	UserId string `json:"userId" gorm:"not null"`
}

func (c *Contact) Cancel() {
	c.Active = false
}

func NewContact(name string, email string, userId string) *Contact {

	return &Contact{
		Name:   name,
		Email:  email,
		Active: true,
		Entity: domain.NewEntity(),
		UserId: userId,
	}
}
