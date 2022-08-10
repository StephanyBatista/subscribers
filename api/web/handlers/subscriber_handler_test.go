package handlers_test

import (
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/helpers/fake"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_campaign_send_post_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/campaigns/any_id/send", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_campaign_send_post_validate_campaign(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/campaigns/any_id/send", nil, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_campaign_send_post_must_campaign_change_status_to_sending(t *testing.T) {
	fake.Build()
	userId := "232cecsd2"
	campaign := createNewCampaign(userId, "teste")
	repository := fake.DI.SubscriberHander.CampaignRepository

	fake.MakeTestHTTP("POST", "/campaigns/"+campaign.ID+"/send", nil, fake.GenerateTokenWithUserId(userId))

	campaignUpdated := repository.GetBy(campaigns.Campaign{Entity: domain.Entity{ID: campaign.ID}})
	assert.Equal(t, campaigns.Sending, campaignUpdated.Status)
}

func Test_campaign_send_post_must_create_subscriber_from_clients(t *testing.T) {
	fake.Build()
	userId := "232cecsd2"
	campaign := createNewCampaign(userId, "teste")
	repository := fake.DI.SubscriberHander.SubscriberRepository
	createNewClient(userId)
	createNewClient(userId)
	amountOfClients := 2

	fake.MakeTestHTTP("POST", "/campaigns/"+campaign.ID+"/send", nil, fake.GenerateTokenWithUserId(userId))
	time.Sleep(1 * time.Second)

	subscribers := repository.List(campaigns.Subscriber{CampaignID: campaign.ID})
	assert.Equal(t, amountOfClients, len(*subscribers))
}

func Test_campaign_send_post_must_try_send_email(t *testing.T) {
	fake.Build()
	userId := "232cecsd2"
	campaign := createNewCampaign(userId, "teste")
	repository := fake.DI.SubscriberHander.SubscriberRepository
	createNewClient(userId)

	fake.MakeTestHTTP("POST", "/campaigns/"+campaign.ID+"/send", nil, fake.GenerateTokenWithUserId(userId))
	time.Sleep(1 * time.Second)

	subscribers := repository.List(campaigns.Subscriber{CampaignID: campaign.ID})

	for _, subscriber := range *subscribers {
		assert.Equal(t, campaigns.NotSent, subscriber.Status)
	}
}
