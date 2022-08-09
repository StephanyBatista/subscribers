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
	"time"

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
	clients := h.ClientRepository.List(clients.Client{UserId: claim.UserId})
	go send(*campaign, *clients, h.SubscriberRepository)

	c.JSON(http.StatusOK, "OK")
}

func send(campaign campaigns.Campaign, clients []clients.Client, repository database.IRepository[campaigns.Subscriber]) {
	for _, client := range clients {
		time.Sleep(2 * time.Second)
		log.Println("Send email to " + client.Email)
		subscriber := campaigns.NewSubscriber(campaign, client.ID, client.Email)
		ok :=
			email.Send(campaign.From, client.Email, campaign.Subject, campaign.Body, campaign.ID, client.ID)
		if ok {
			subscriber.Sent()
		} else {
			subscriber.NotSent()
		}
		repository.Save(subscriber)
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
