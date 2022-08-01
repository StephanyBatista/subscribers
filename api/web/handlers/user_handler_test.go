package handlers_test

import (
	"io/ioutil"
	"net/http"
	"subscribers/domain/users"
	"subscribers/helpers/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFieldsRequiredPost(t *testing.T) {
	fake.Build()
	w := fake.MakeTestHTTP("POST", "/users", nil, "")

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Name' is required")
	assert.Contains(t, responseString, "'Email' is required")
	assert.Contains(t, responseString, "'Password' is required")
}

func TestValidateInvalidEmailPost(t *testing.T) {
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    "invalid",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Email' is invalid")
}

func TestValidateEmailAlreadySavedPost(t *testing.T) {
	fake.Build()
	userSaved, _ := users.NewUser(users.CreationRequest{Name: "Teste", Email: "teste@teste.com.br", Password: "password123"})
	fake.DB.Create(&userSaved)
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    userSaved.Email,
		Password: "35 million",
	}
	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "Email already saved")
}

func TestSaveNewUserPost(t *testing.T) {
	fake.Build()
	newUser := users.CreationRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := fake.MakeTestHTTP("POST", "/users", newUser, "")

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}
