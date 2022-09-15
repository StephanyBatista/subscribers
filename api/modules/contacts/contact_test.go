package contacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_new_contact(t *testing.T) {

	nameExpected := "test"
	emailExpected := "test@test.com.br"
	userId := "23121"
	contact, _ := NewContact(nameExpected, emailExpected, userId)

	assert.Equal(t, nameExpected, contact.Name)
	assert.Equal(t, emailExpected, contact.Email)
	assert.Equal(t, userId, contact.UserId)
}

func Test_must_validate_fields_when_create_contact(t *testing.T) {

	_, errs := NewContact("", "", "")

	assert.Equal(t, "'Name' is required", errs[0].Error())
	assert.Equal(t, "'Email' is required", errs[1].Error())
	assert.Equal(t, "'UserId' is required", errs[2].Error())
}

func Test_must_validate_email_when_create_contact(t *testing.T) {

	_, errs := NewContact("henrique", "email_invalid", "2323")

	assert.Equal(t, "'Email' invalid", errs[0].Error())
}

func Test_must_cancel_contact(t *testing.T) {

	contact := Contact{Active: true}

	contact.Cancel()

	assert.False(t, contact.Active)
}
