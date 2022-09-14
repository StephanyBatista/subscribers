package files

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"subscribers/utils/web/middlewares"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{}
	secured := router.Group("").Use(middlewares.Auth())
	{
		secured.POST("/files", handler.Post)
	}
}
