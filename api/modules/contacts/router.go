package contacts

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"subscribers/modules/users"
	"subscribers/utils/web/middlewares"
)

func ApplyRouter(router *gin.Engine, db *sql.DB) {
	handler := Handler{
		UserRepository:    users.Repository{DB: db},
		ContactRepository: Repository{DB: db},
	}
	router.PATCH("/contacts/:id/cancel", handler.Cancel)
	secured := router.Group("").Use(middlewares.Auth())
	{
		secured.POST("/contacts", handler.Post)
		secured.GET("/contacts", handler.GetAll)
		secured.GET("/contacts/:id", handler.GetById)
	}
}
