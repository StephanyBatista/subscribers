package handlers

import (
	"net/http"
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
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewErrorReponse("'file' is required"))
		return
	}

	defer file.Close()
	url, ok := storage.Upload(file, header.Filename)
	if !ok {
		c.JSON(http.StatusBadRequest, web.NewInternalError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"url": url})
}
