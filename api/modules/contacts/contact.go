package contacts

import (
	"errors"
	"net/mail"
	"time"

	"github.com/rs/xid"
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

func NewContact(name string, email string, userId string) (Contact, []error) {

	errs := validate(name, email, userId)
	if len(errs) > 0 {
		return Contact{}, errs
	}

	return Contact{
		Id:        xid.New().String(),
		CreatedAt: time.Now().UTC(),
		Name:      name,
		Email:     email,
		Active:    true,
		UserId:    userId,
	}, nil
}

func validate(name, email, userId string) []error {
	var errs []error
	if name == "" {
		errs = append(errs, errors.New("'Name' is required"))
	}
	if email == "" {
		errs = append(errs, errors.New("'Email' is required"))
	} else if _, err := mail.ParseAddress(email); err != nil {
		errs = append(errs, errors.New("'Email' invalid"))
	}
	if userId == "" {
		errs = append(errs, errors.New("'UserId' is required"))
	}
	return errs
}
