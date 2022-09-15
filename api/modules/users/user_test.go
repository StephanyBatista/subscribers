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

func Test_create_new_user_must_validate_fields(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	_, errs := NewUser("", "", "")

	assert.Equal(t, "'Name' is required", errs[0].Error())
	assert.Equal(t, "'Email' is required", errs[1].Error())
	assert.Equal(t, "'Password' is required", errs[2].Error())
}

func Test_create_new_user_must_validate_name_size(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	_, errs := NewUser("hen", "", "")

	assert.Equal(t, "'Name' invalid size, min 4", errs[0].Error())
}

func Test_create_new_user_must_validate_email(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	_, errs := NewUser("henrique", "email_invalid", "")

	assert.Equal(t, "'Email' invalid", errs[0].Error())
}

func Test_create_new_user_must_validate_password_size(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	_, errs := NewUser("henrique", "email@email.com", "12345")

	assert.Equal(t, "'Password' invalid size, min 6", errs[0].Error())
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

func Test_must_validate_fields_when_change_password(t *testing.T) {
	os.Setenv("sub_salt_hash", "5")

	user := User{}

	errs := user.ChangePassword("", "")

	assert.Equal(t, "'OldPassword' is required", errs[0].Error())
	assert.Equal(t, "'NewPassword' is required", errs[1].Error())
}
