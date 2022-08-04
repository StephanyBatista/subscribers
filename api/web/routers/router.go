package routers

import (
	"subscribers/helpers"
	"subscribers/web/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter(di *helpers.DI) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		//ExposeHeaders: []string{"Content-Length"},
	}))
	r.GET("/healthcheck", di.HealthCheckHandler.Get)
	r.POST("/token", di.TokenHandler.Post)
	r.POST("/users", di.UserHandler.Post)
	secured := r.Group("").Use(middlewares.Auth())
	{
		secured.GET("/users/info", di.UserHandler.GetInfo)
		secured.POST("/campaigns", di.CampaignHandler.Post)
		secured.GET("/campaigns/:id", di.CampaignHandler.GetById)
		secured.GET("/campaigns", di.CampaignHandler.GetAll)
	}
	
	return r
}
