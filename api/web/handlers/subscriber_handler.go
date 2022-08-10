package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/domain/clients"
	"subscribers/infra/database"
	"subscribers/infra/email"
	"subscribers/web"
	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type SubscriberHandler struct {
	CampaignRepository   database.IRepository[campaigns.Campaign]
	SubscriberRepository database.IRepository[campaigns.Subscriber]
	ClientRepository     database.IRepository[clients.Client]
}

func (h *SubscriberHandler) Post(c *gin.Context) {

	campaignId := c.Param("campaignID")
	campaign := h.CampaignRepository.GetBy(campaigns.Campaign{Entity: domain.Entity{ID: campaignId}})
	if campaign == nil {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))
	campaign.Sending()
	h.CampaignRepository.Save(campaign)

	c.JSON(http.StatusOK, "OK")

	go processSubscribers(claim.UserId, *campaign, h.ClientRepository, h.SubscriberRepository)
}

func processSubscribers(
	userId string,
	campaign campaigns.Campaign,
	clientRepository database.IRepository[clients.Client],
	subscribersRepository database.IRepository[campaigns.Subscriber]) {

	clients := clientRepository.List(clients.Client{UserId: userId})
	if clients == nil {
		return
	}

	subscribers := make([]campaigns.Subscriber, len(*clients))
	for index, client := range *clients {
		subscriber := campaigns.NewSubscriber(campaign, client.ID, client.Email)
		subscribersRepository.Create(subscriber)
		subscribers[index] = *subscriber
	}

	for _, subscriber := range subscribers {
		log.Println("Send email to " + subscriber.Email)
		ok :=
			email.Send(campaign.From, subscriber.Email, campaign.Subject, campaign.Body, campaign.ID, subscriber.ClientID)
		if ok {
			subscriber.Sent()
		} else {
			subscriber.NotSent()
		}
		subscribersRepository.Save(&subscriber)
	}
}

func (h *SubscriberHandler) GetRead(c *gin.Context) {

	subscriberId := c.Param("id")
	entity := h.SubscriberRepository.GetBy(campaigns.Subscriber{Entity: domain.Entity{ID: subscriberId}})
	if entity != nil {
		entity.Read()
		h.SubscriberRepository.Save(entity)
	}
	c.String(http.StatusOK, "img")
}
