package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribers/commun/storage"
	"subscribers/commun/web"
)

type Handler struct {
}

func (h *Handler) Post(c *gin.Context) {
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
