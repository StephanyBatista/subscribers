package campaigns

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"subscribers/web/middlewares"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{
		CampaignRepository: Repository{DB: db},
	}
	secured := router.Group("").Use(middlewares.Auth())
	{
		secured.POST("/campaigns", handler.Post)
		secured.GET("/campaigns/:id", handler.GetById)
		secured.GET("/campaigns", handler.GetAll)
		secured.POST("/campaigns/:campaignID/ready", handler.Ready)
	}
}
