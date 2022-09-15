package healtchcheck

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{
		Db: db,
	}
	router.GET("/healthcheck", handler.Get)
}
