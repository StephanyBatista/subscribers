package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"subscribers/commun/web/middlewares"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{UserRepository: Repository{DB: db}}
	router.POST("/users", handler.Post)
	router.POST("/token", handler.Token)
	secured := router.Group("").Use(middlewares.Auth())
	{
		secured.GET("/users/info", handler.GetInfo)
		secured.PATCH("/users/changepassword", handler.ChangePassword)
	}
}
