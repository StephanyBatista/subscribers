package healtchcheck

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db *sql.DB
}

func (h *Handler) Get(c *gin.Context) {
	err := h.Db.Ping()
	database := false
	if err == nil {
		database = true
	}

	c.JSON(http.StatusOK, gin.H{"database": database})
}
