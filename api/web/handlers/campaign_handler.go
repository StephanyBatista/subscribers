package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/domain/clients"
	"subscribers/infra/database"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	CampaignRepository   database.IRepository[campaigns.Campaign]
	SubscriberRepository database.IRepository[campaigns.Subscriber]
	ClientRepository     database.IRepository[clients.Client]
}

func (h *CampaignHandler) Post(c *gin.Context) {
	var body CampaignRequest
	c.BindJSON(&body)
	errs := domain.Validate(body)
	if errs != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity := campaigns.NewCampaign(body.Name, body.From, body.Subject, body.Body, claim.UserId, claim.UserName)
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

	entity := h.CampaignRepository.GetBy(campaigns.Campaign{Entity: domain.Entity{ID: id}})
	if entity == nil || claim.UserId != entity.CreatedBy.Id {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (h *CampaignHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entities := h.CampaignRepository.List(campaigns.Campaign{CreatedBy: domain.UserValue{Id: claim.UserId}})
	if entities == nil {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entities)
}
