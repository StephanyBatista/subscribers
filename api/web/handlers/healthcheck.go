package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheck struct {
	Db *gorm.DB
}

func (h *HealthCheck) Get(c *gin.Context) {
	_, err := h.Db.Raw("SELECT backend_start as date FROM pg_stat_activity limit 1").Rows()
	database := false
	if err == nil {
		database = true
	}

	c.JSON(http.StatusOK, gin.H{"database": database})
}

func NewHealthCheck(db *gorm.DB) *HealthCheck {
	return &HealthCheck{Db: db}
}
