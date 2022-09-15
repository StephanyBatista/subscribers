package campaigns

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"subscribers/utils/web"
	"subscribers/utils/web/auth"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CampaignRepository Repository
	Session            *session.Session
}

func (h *Handler) Post(c *gin.Context) {
	var body CreateNewCampaign
	c.BindJSON(&body)

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity, errs := NewCampaign(body.Name, body.From, body.Subject, body.Body, claim.UserId)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}
	err := h.CampaignRepository.Create(entity)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.Id})
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity, _ := h.CampaignRepository.GetBy(id)
	if entity.Id == "" || claim.UserId != entity.UserId {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	response := CampaignResponse{
		ID:      entity.Id,
		Name:    entity.Name,
		Status:  entity.Status,
		From:    entity.From,
		Subject: entity.Subject,
		Body:    entity.Body,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetEmailsReport(c *gin.Context) {
	id := c.Param("id")

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity, err := h.CampaignRepository.GetBy(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	if entity.Id == "" || claim.UserId != entity.UserId {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}

	report, err := h.CampaignRepository.GetEmailsReport(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *Handler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities, _ := h.CampaignRepository.ListBy(claim.UserId)
	if entities == nil || len(entities) == 0 {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entities)
}

func (h *Handler) Ready(c *gin.Context) {

	campaignId := c.Param("campaignID")
	campaign, _ := h.CampaignRepository.GetBy(campaignId)
	if campaign.Id == "" {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	} else if campaign.Status != Draft {
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
