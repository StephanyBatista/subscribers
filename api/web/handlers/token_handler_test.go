package handlers_test

import (
	"io/ioutil"
	"net/http"
	"subscribers/domain/user"
	"subscribers/helpers"
	"subscribers/web/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenPostValidateFieldsRequired(t *testing.T) {
	di := helpers.NewFakeDI()
	w := helpers.CreateHTTPTest("POST", "/token", di.TokenHandler.Post, nil)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Email' is required")
	assert.Contains(t, responseString, "'Password' is required")
}

func TestTokenPostUserNotFound(t *testing.T) {
	body := handlers.LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}

	w := helpers.CreateHTTPTest("POST", "/token", di.TokenHandler.Post, body)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "User not found")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestTokenPostGenerateJwt(t *testing.T) {
	di := helpers.NewFakeDI()
	body := handlers.LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}
	user, _ := user.NewUser("test", body.Email, body.Password)
	di.DB.Create(user)

	w := helpers.CreateHTTPTest("POST", "/token", di.TokenHandler.Post, body)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "token")
	assert.Contains(t, responseString, "expiresAt")
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
