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

var di *helpers.DI = helpers.NewFakeDI()

func TestValidateFieldsRequiredPost(t *testing.T) {
	w := helpers.CreateHTTPTest("POST", "/users", di.UserHandler.Post, nil)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Name' is required")
	assert.Contains(t, responseString, "'Email' is required")
	assert.Contains(t, responseString, "'Password' is required")
}

func TestValidateInvalidEmailPost(t *testing.T) {
	newUser := handlers.UserCreationRequest{
		Name:     "Demo",
		Email:    "invalid",
		Password: "35 million",
	}

	w := helpers.CreateHTTPTest("POST", "/users", di.UserHandler.Post, newUser)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Email' is invalid")
}

func TestValidateEmailAlreadySavedPost(t *testing.T) {
	userSaved, _ := user.NewUser("Teste", "teste@teste.com.br", "password123")
	di.DB.Create(&userSaved)
	newUser := handlers.UserCreationRequest{
		Name:     "Demo",
		Email:    userSaved.Email,
		Password: "35 million",
	}
	w := helpers.CreateHTTPTest("POST", "/users", di.UserHandler.Post, newUser)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "Email already saved")
}

func TestSaveNewUserPost(t *testing.T) {
	newUser := handlers.UserCreationRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := helpers.CreateHTTPTest("POST", "/users", di.UserHandler.Post, newUser)

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
