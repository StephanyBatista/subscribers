package campaigns

import (
	"database/sql"
	"subscribers/commun/queue"
	"subscribers/commun/web/middlewares"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.Engine, db *sql.DB, queue queue.IQueue) {
	handler := Handler{
		CampaignRepository: Repository{DB: db},
		Queue:              queue,
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
