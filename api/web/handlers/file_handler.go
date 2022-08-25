package handlers

import (
	"errors"
	"net/http"
	"subscribers/domain"
	"subscribers/domain/campaigns"
	"subscribers/infra/database"
	"subscribers/infra/storage"
	"subscribers/web"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	CampaignRepository database.IRepository[campaigns.Campaign]
}

func (h *FileHandler) Post(c *gin.Context) {

	var errs []error
	if c.Request.FormValue("kind") != "campaign" {
		errs = append(errs, errors.New("'kind' is required"))
	}

	if c.Request.FormValue("keyId") == "" {
		errs = append(errs, errors.New("'keyId' is required"))
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		errs = append(errs, errors.New("'file' is required"))
	}

	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, web.NewErrorsReponse(errs))
		return
	}

	defer file.Close()
	url, ok := storage.Upload(file, header.Filename)
	if !ok {
		c.JSON(http.StatusBadRequest, web.NewInternalError())
		return
	}
	keyId := c.Request.FormValue("keyId")
	kind := c.Request.FormValue("kind")

	if kind == "campaign" {
		campaign := h.CampaignRepository.GetBy(campaigns.Campaign{Entity: domain.Entity{ID: keyId}})
		if campaign == nil {
			c.JSON(http.StatusNotFound, web.NewErrorReponse("campaign not found"))
			return
		}
		campaign.AttachmentURL = url
		h.CampaignRepository.Save(campaign)
	}
	c.JSON(http.StatusCreated, gin.H{"url": url})
}
