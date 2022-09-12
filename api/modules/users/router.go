package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{UserRepository: Repository{DB: db}}
	router.POST("/users", handler.Post)
	router.POST("/token", handler.Token)
}
