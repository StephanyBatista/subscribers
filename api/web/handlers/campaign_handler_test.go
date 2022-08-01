package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/helpers/fake"
	"subscribers/web/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCampaignPostValidateToken(t *testing.T) {
	fake.Build()
	w := fake.MakeTestHTTP("POST", "/campaigns", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCampaignPostValidateFields(t *testing.T) {
	fake.Build()
	token, _, _ := auth.GenerateJWT("xpto", "teste@teste.com.br", "test")

	w := fake.MakeTestHTTP("POST", "/campaigns", nil, token)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, "'Name' is required")
	assert.Contains(t, responseString, "'Active' is required")
}

func TestCampaignPostSaveNewCampaign(t *testing.T) {
	fake.Build()
	token, _, _ := auth.GenerateJWT("xpto", "teste@teste.com.br", "test")
	body := campaigns.CreationRequest{Name: "teste 1", Active: true}

	w := fake.MakeTestHTTP("POST", "/campaigns", body, token)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCampaignGetCampaignById(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	token, _, _ := auth.GenerateJWT(userValue.Id, "teste@teste.com.br", userValue.Name)
	request := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	entity, _ := campaigns.NewCampaign(request, userValue)
	fake.DB.Create(&entity)

	w := fake.MakeTestHTTP("GET", "/campaigns/"+entity.ID, entity, token)

	responseData, _ := ioutil.ReadAll(w.Body)
	responseString := string(responseData)
	assert.Contains(t, responseString, request.Name)
	assert.Contains(t, responseString, request.Description)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCampaignGetAllCampaignOfUser(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	token, _, _ := auth.GenerateJWT(userValue.Id, "teste@teste.com.br", userValue.Name)
	request1 := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	entity1, _ := campaigns.NewCampaign(request1, userValue)
	fake.DB.Create(&entity1)
	request2 := campaigns.CreationRequest{Name: "teste 2", Description: "AAx tere", Active: true}
	entity2, _ := campaigns.NewCampaign(request2, userValue)
	fake.DB.Create(&entity2)
	requestOfOtherUser := campaigns.CreationRequest{Name: "teste 2", Description: "AAx tere", Active: true}
	entity3, _ := campaigns.NewCampaign(requestOfOtherUser, domain.UserValue{Id: "Fs156", Name: "outher user"})
	fake.DB.Create(&entity3)
	amountOfCampaignsExpected := 2

	w := fake.MakeTestHTTP("GET", "/campaigns", nil, token)

	responseData, _ := ioutil.ReadAll(w.Body)
	campaignsResult := []campaigns.Campaign{}
	json.Unmarshal([]byte(responseData), &campaignsResult)
	assert.Equal(t, amountOfCampaignsExpected, len(campaignsResult))
}
