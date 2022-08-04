package handlers_test

import (
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/domain/campaigns/clients"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientPostValidateFields(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/clients", nil, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'CampaignId' is required")
}

func TestClientPostValidateCampaign(t *testing.T) {
	fake.Build()
	body := clients.CreationRequest{
		Name:       "Teste",
		Email:      "teste@teste.com.br",
		CampaignId: "INVALID_CAMPAIGN",
	}

	w := fake.MakeTestHTTP("POST", "/clients", body, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "Campaign not found")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestClientPostSaveNewClient(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	request := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	campaign, _ := campaigns.NewCampaign(request, userValue)
	fake.DB.Create(&campaign)
	body := clients.CreationRequest{
		Name:       "Teste",
		Email:      "teste@teste.com.br",
		CampaignId: campaign.ID,
	}

	w := fake.MakeTestHTTP("POST", "/clients", body, "")

	fake.DB.Where(campaigns.Campaign{Entity: &domain.Entity{ID: campaign.ID}}).Preload("Clients").First(&campaign)
	assert.True(t, campaign.HasClient(body.Email))
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestClientPostNotSaveTheSameClientOnCampaign(t *testing.T) {
	fake.Build()
	userValue := domain.UserValue{Id: "xpto", Name: "test"}
	request := campaigns.CreationRequest{Name: "teste 1", Description: "AAx tere", Active: true}
	campaign, _ := campaigns.NewCampaign(request, userValue)
	fake.DB.Create(&campaign)
	body := clients.CreationRequest{
		Name:       "Teste",
		Email:      "teste@teste.com.br",
		CampaignId: campaign.ID,
	}

	fake.MakeTestHTTP("POST", "/clients", body, "")
	fake.MakeTestHTTP("POST", "/clients", body, "")

	fake.DB.Where(campaigns.Campaign{Entity: &domain.Entity{ID: campaign.ID}}).Preload("Clients").First(&campaign)
	assert.True(t, len(campaign.Clients) == 1)

}
