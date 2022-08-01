package handlers_test

import (
	"net/http"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserPostValidateFieldsRequired(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/users", nil, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'Password' is required")
}

func TestUserPostValidateInvalidEmail(t *testing.T) {
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    "invalid",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Email' is invalid")
}

func TestUserPostValidateWhenEmailIsBeingUsed(t *testing.T) {
	fake.Build()
	userSaved, _ :=
		users.NewUser(users.CreationRequest{Name: "Teste", Email: "teste@teste.com.br", Password: "password123"})
	fake.DB.Create(&userSaved)
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    userSaved.Email,
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "Email already saved")
}

func TestUserPostSaveNewUser(t *testing.T) {
	fake.Build()
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
