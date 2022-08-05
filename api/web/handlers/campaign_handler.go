package handlers

import (
	"log"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/web"

	"subscribers/web/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CampaignHandler struct {
	Db *gorm.DB
}

func (h *CampaignHandler) Post(c *gin.Context) {
	var body campaigns.CreationRequest
	c.BindJSON(&body)

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	entity, errs := campaigns.NewCampaign(body, domain.UserValue{Id: claim.UserId, Name: claim.UserName})
	if errs != nil {
		log.Println(errs)
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	result := h.Db.Create(&entity)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": entity.ID})
}

func (h *CampaignHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	var entity campaigns.Campaign
	result := h.Db.Where(campaigns.Campaign{Entity: domain.Entity{ID: id}}).Preload("Clients").FirstOrInit(&entity)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	} else if entity.IDIsNull() || claim.UserId != entity.CreatedBy.Id {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (h *CampaignHandler) GetAll(c *gin.Context) {
	claim, _ := auth.GetClaimFromToken(c.GetHeader("Authorization"))

	var entities []campaigns.Campaign
	result := h.Db.Where(campaigns.Campaign{CreatedBy: domain.UserValue{Id: claim.UserId}}).Find(&entities)
	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusInternalServerError, web.NewInternalError())
		return
	} else if len(entities) == 0 {
		log.Println("Campaign not found")
		c.JSON(http.StatusNotFound, web.NewErrorReponse("Not found"))
		return
	}
	c.JSON(http.StatusOK, entities)
}
