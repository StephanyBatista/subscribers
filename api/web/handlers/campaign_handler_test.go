package handlers_test

import (
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/helpers"
	"subscribers/helpers/fake"
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

	w := fake.MakeTestHTTP("POST", "/campaigns", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Active' is required")
}

func TestCampaignPostSaveNewCampaign(t *testing.T) {
	fake.Build()
	body := campaigns.CreationRequest{Name: "teste 1", Active: true}

	w := fake.MakeTestHTTP("POST", "/campaigns", body, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCampaignGetCampaignById(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	request := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	entity, _ := campaigns.NewCampaign(request, userValue)
	fake.DB.Create(&entity)

	w := fake.MakeTestHTTP("GET", "/campaigns/"+entity.ID, entity, fake.GenerateTokenWithUserId(userValue.Id))

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, request.Name)
	assert.Contains(t, response, request.Description)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCampaignGetByIdNotFound(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/campaigns/id_invalid", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestCampaignGetAllCampaignOfUser(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	request1 := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	entity1, _ := campaigns.NewCampaign(request1, userValue)
	fake.DB.Create(&entity1)
	request2 := campaigns.CreationRequest{Name: "teste 2", Description: "AAx tere", Active: true}
	entity2, _ := campaigns.NewCampaign(request2, userValue)
	fake.DB.Create(&entity2)
	requestOfOtherUser := campaigns.CreationRequest{Name: "teste 2", Description: "AAx tere", Active: true}
	entityOfAnotherUser, _ := campaigns.NewCampaign(requestOfOtherUser, domain.UserValue{Id: "Fs156", Name: "outher user"})
	fake.DB.Create(&entityOfAnotherUser)
	amountOfCampaignsExpectedOfUser := 2

	w := fake.MakeTestHTTP("GET", "/campaigns", nil, fake.GenerateTokenWithUserId(userValue.Id))

	campaignsOfUser := helpers.BufferToObj[[]campaigns.Campaign](w.Body)
	assert.Equal(t, amountOfCampaignsExpectedOfUser, len(campaignsOfUser))
}

func TestCampaignGetAllNotFound(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/campaigns", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}
