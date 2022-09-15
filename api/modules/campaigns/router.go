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
	secured := router.Group("campaigns").Use(middlewares.Auth())
	{
		secured.POST("", handler.Post)
		secured.POST("/:id/ready", handler.Ready)
		secured.GET("", handler.GetAll)
		secured.GET("/:id", handler.GetById)
		secured.GET("/:id/emailsreport", handler.GetEmailsReport)
	}
}
