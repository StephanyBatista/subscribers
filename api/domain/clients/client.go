package clients

import "subscribers/domain"

type Client struct {
	domain.Entity
	Name   string `json:"name" gorm:"not null;size:100;"`
	Email  string `json:"email" gorm:"not null;size:100;"`
	Active bool   `json:"active" gorm:"not null"`
	UserId string `json:"userId" gorm:"not null"`
}

func NewClient(name string, email string, userId string) *Client {

	client := &Client{
		Name:   name,
		Email:  email,
		Active: true,
		Entity: domain.NewEntity(),
		UserId: userId,
	}

	return client
}
