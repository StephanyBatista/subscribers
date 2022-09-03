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
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders: []string{"Content-Length"},
	}))
	r.GET("/healthcheck", di.HealthCheckHandler.Get)
	r.POST("/token", di.TokenHandler.Post)
	r.POST("/users", di.UserHandler.Post)
	r.GET("/subscribers/:id/read", di.SubscriberHander.GetRead)

	secured := r.Group("").Use(middlewares.Auth())
	{
		secured.GET("/users/info", di.UserHandler.GetInfo)
		secured.PATCH("/users/changepassword", di.UserHandler.ChangePassword)
		secured.POST("/campaigns", di.CampaignHandler.Post)
		secured.GET("/campaigns/:id", di.CampaignHandler.GetById)
		secured.GET("/campaigns", di.CampaignHandler.GetAll)
		secured.POST("/campaigns/:campaignID/send", di.SubscriberHander.Post)
		secured.POST("/contacts", di.ContactHandler.Post)
		secured.GET("/contacts", di.ContactHandler.GetAll)
		secured.GET("/contacts/:id", di.ContactHandler.GetById)
		secured.POST("/files", di.FileHandler.Post)
	}

	return r
}
