package handlers_test

import (
	"io/ioutil"
	"net/http"
	"subscribers/domain/users"
	"subscribers/helpers/fake"
	"subscribers/web/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenPostValidateFieldsRequired(t *testing.T) {
	fake.Build()
	w := fake.MakeTestHTTP("POST", "/token", nil, "")

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Email' is required")
	assert.Contains(t, responseString, "'Password' is required")
}

func TestTokenPostUserNotFound(t *testing.T) {
	fake.Build()
	body := handlers.LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/token", body, "")

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "User not found")
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

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "token")
	assert.Contains(t, responseString, "expiresAt")
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
