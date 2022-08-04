package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/domain/campaigns/clients"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientHandler struct {
	Db *gorm.DB
}

func (h *ClientHandler) Post(c *gin.Context) {
	var body clients.CreationRequest
	c.BindJSON(&body)

	claim, claimOk := auth.GetClaimFromToken(c.GetHeader("Authorization"))
	var user *domain.UserValue = nil

	if claimOk {
		user = &domain.UserValue{Id: claim.UserId, Name: claim.UserName}
	}

	entity, errs := clients.NewClient(body, user)
	if errs != nil {
		log.Println(errs)
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	var campaign campaigns.Campaign
	result := h.Db.Where(campaigns.Campaign{Entity: &domain.Entity{ID: body.CampaignId}}).Preload("Clients").FirstOrInit(&campaign)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Campaign not found"))
		return
	}
	var clientSaved clients.Client
	result = h.Db.Where(clients.Client{Email: body.Email}).FirstOrInit(&clientSaved)
	if result.RowsAffected > 0 {
		if campaign.HasClient(clientSaved.Email) {
			c.JSON(http.StatusCreated, gin.H{"id": clientSaved.ID})
			return
		}
		entity = &clientSaved
	} else {
		result = h.Db.Create(&entity)
		if result.Error != nil {
			log.Println("test", result.Error)
			c.JSON(http.StatusInternalServerError, web.NewInternalError())
			return
		}
	}

	campaign.AddClient(entity)
	result = h.Db.Save(&campaign)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}
