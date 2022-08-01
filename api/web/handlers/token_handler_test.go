package handlers_test

import (
	"net/http"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"subscribers/web/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenPostValidateFieldsRequired(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/token", nil, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'Password' is required")
}

func TestTokenPostUserNotFound(t *testing.T) {
	fake.Build()
	body := handlers.LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/token", body, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "User not found")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestTokenPostGenerateJwt(t *testing.T) {
	fake.Build()
	body := handlers.LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}
	user, _ := users.NewUser(users.CreationRequest{Name: "Teste", Email: body.Email, Password: body.Password})
	fake.DB.Create(user)

	w := fake.MakeTestHTTP("POST", "/token", body, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "token")
	assert.Contains(t, response, "expiresAt")
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
