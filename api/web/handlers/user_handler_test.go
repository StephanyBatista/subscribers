package handlers_test

import (
	"net/http"
	"os"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"subscribers/web/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_user_post_validate_fields_required(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/users", nil, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'Password' is required")
}

func Test_user_post_validate_invalid_email(t *testing.T) {
	fake.Build()
	newUser := handlers.UserRequest{
		Name:     "Demo",
		Email:    "invalid",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Email' is invalid")
}

func Test_user_post_validate_when_email_is_being_used(t *testing.T) {
	fake.Build()
	userSaved, _ :=
		users.NewUser("Teste", "teste@teste.com.br", "password123")
	fake.DB.Create(&userSaved)
	newUser := handlers.UserRequest{
		Name:     "Demo",
		Email:    userSaved.Email,
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "Email already saved")
}

func Test_user_post_save_new_user(t *testing.T) {
	fake.Build()
	newUser := handlers.UserRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}

func Test_user_post_show_error_when_try_to_generate_password_hash(t *testing.T) {
	fake.Build()
	invalidSalt := "23323232323"
	os.Setenv("sub_salt_hash", invalidSalt)
	newUser := handlers.UserRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "1",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "error to generate password")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func Test_user_post_show_error_when_not_create(t *testing.T) {
	fake.Build()
	mock := &fake.RepositoryMock[users.User]{
		ReturnsCreate: false,
	}
	fake.DI.UserHandler.UserRepository = mock
	newUser := handlers.UserRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
