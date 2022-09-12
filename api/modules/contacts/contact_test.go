package contacts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_new_contact(t *testing.T) {

	nameExpected := "test"
	emailExpected := "test@test.com.br"
	userId := "23121"
	contact := NewContact(nameExpected, emailExpected, userId)

	assert.Equal(t, nameExpected, contact.Name)
	assert.Equal(t, emailExpected, contact.Email)
	assert.Equal(t, userId, contact.UserId)
}

func Test_must_cancel_contact(t *testing.T) {

	contact := Contact{Active: true}

	contact.Cancel()

	assert.False(t, contact.Active)
}
