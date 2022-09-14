package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/infra/database"
	campaigns2 "subscribers/modules/campaigns"
	"subscribers/modules/contacts"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	CampaignRepository   database.IRepository[campaigns2.Campaign]
	SubscriberRepository database.IRepository[campaigns.Subscriber]
	ContactRepository    database.IRepository[contacts.Contact]
	Session              *session.Session
}

func (h *CampaignHandler) Post(c *gin.Context) {
	var body CampaignRequest
	c.BindJSON(&body)

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := campaigns2.NewCampaign(body.Name, body.From, body.Subject, body.Body, claim.UserId, claim.UserName)
	ok := h.CampaignRepository.Create(entity)
	if !ok {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}

func (h *CampaignHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := h.CampaignRepository.GetBy(campaigns2.Campaign{Entity: domain.Entity{ID: id}})
	if entity == nil || claim.UserId != entity.CreatedBy.Id {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	response := CampaignResponse{
		ID:            entity.ID,
		Name:          entity.Name,
		Status:        entity.Status,
		From:          entity.From,
		Subject:       entity.Subject,
		Body:          entity.Body,
		AttachmentURL: entity.AttachmentURL,
	}

	subscribers := h.SubscriberRepository.List(campaigns.Subscriber{CampaignID: entity.ID})
	if subscribers != nil {
		for _, subscriber := range *subscribers {
			response.BaseOfSubscribers++
			if subscriber.Status == campaigns2.Sent {
				response.TotalSent++
			} else if subscriber.Status == campaigns2.Read {
				response.TotalRead++
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities := h.CampaignRepository.List(campaigns2.Campaign{CreatedBy: domain.UserValue{Id: claim.UserId}})
	if entities == nil {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entities)
}

func (h *CampaignHandler) Send(c *gin.Context) {

	campaignId := c.Param("campaignID")
	campaign := h.CampaignRepository.GetBy(campaigns2.Campaign{Entity: domain.Entity{ID: campaignId}})
	if campaign == nil {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	} else if campaign.Status != campaigns2.Draft {
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("Campaigns is with different status"))
		return
	}

	sqsClient := sqs.New(h.Session)

	messageBody := fmt.Sprintf(`{"Id": "%s"}`, campaignId)
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(os.Getenv("AWS_URL_QUEUE_CAMPAIGN_READY")),
		MessageBody: aws.String(messageBody),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, web.NewErrorReponse("Error to save on queue"))
		return
	}

	campaign.Ready()
	h.CampaignRepository.Save(campaign)

	c.JSON(http.StatusOK, "OK")
}
