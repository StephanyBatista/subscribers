package users

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_new_user(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")
	nameExpected := "Henrique"
	emailExpected := "henrique@gmail.com"
	password := "123456"

	user, _ := NewUser(nameExpected, emailExpected, password)

	assert.Equal(t, nameExpected, user.Name)
	assert.Equal(t, emailExpected, user.Email)
	assert.NotEmpty(t, user.PasswordHash)
	assert.NotEmpty(t, user.Id)
}

func Test_when_create_new_user_the_date_created_at_must_be_setted(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	user, _ := NewUser("Henrique", "henrique@gmail.com", "123456")

	assert.NotEmpty(t, user.CreatedAt)
}

func Test_new_password_must_be_equals_hash(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")
	password := "123456"

	user, _ := NewUser("Henrique", "henrique@gmail.com", password)

	assert.True(t, user.CheckPassword(password))
}

func Test_must_change_password(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")
	oldPassword := "123456"
	newPassword := "test123"
	user, _ := NewUser("Henrique", "henrique@gmail.com", oldPassword)

	user.ChangePassword(oldPassword, newPassword)

	assert.True(t, user.CheckPassword(newPassword))
}
