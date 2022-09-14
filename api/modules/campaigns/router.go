package campaigns

import (
	"database/sql"
	"subscribers/utils/web/middlewares"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.Engine, db *sql.DB, session *session.Session) {
	handler := Handler{
		CampaignRepository: Repository{DB: db},
		Session:            session,
	}
	secured := router.Group("").Use(middlewares.Auth())
	{
		secured.POST("/campaigns", handler.Post)
		secured.GET("/campaigns/:id", handler.GetById)
		secured.GET("/campaigns", handler.GetAll)
		secured.POST("/campaigns/:campaignID/ready", handler.Ready)
	}
}
