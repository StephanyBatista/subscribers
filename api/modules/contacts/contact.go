package contacts

import (
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Id        string    `json:"id" `
	Name      string    `json:"name" `
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *Contact) Cancel() {
	c.Active = false
}

func NewContact(name string, email string, userId string) Contact {

	return Contact{
		Id:        xid.New().String(),
		CreatedAt: time.Now().UTC(),
		Name:      name,
		Email:     email,
		Active:    true,
		UserId:    userId,
	}
}
