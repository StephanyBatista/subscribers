package helpers

import (
	"subscribers/domain/campaigns"
	"subscribers/domain/clients"
	"subscribers/domain/users"
	"subscribers/infra/database"
	"subscribers/web/handlers"

	"gorm.io/gorm"
)

type DI struct {
	DB                 *gorm.DB
	TokenHandler       *handlers.TokenHandler
	UserHandler        *handlers.UserHandler
	HealthCheckHandler *handlers.HealthCheckHandler
	CampaignHandler    *handlers.CampaignHandler
	SubscriberHander   *handlers.SubscriberHandler
	ClientHandler      *handlers.ClientHandler
}

func NewDI() *DI {
	di := &DI{}
	db := database.CreateConnection()
	di.DB = db
	di.TokenHandler = &handlers.TokenHandler{Db: db}

	di.UserHandler = &handlers.UserHandler{
		UserRepository: &database.Repository[users.User]{DB: db},
	}
	di.HealthCheckHandler = &handlers.HealthCheckHandler{Db: db}
	di.CampaignHandler = &handlers.CampaignHandler{
		CampaignRepository:   &database.Repository[campaigns.Campaign]{DB: db},
		SubscriberRepository: &database.Repository[campaigns.Subscriber]{DB: db},
		ClientRepository:     &database.Repository[clients.Client]{DB: db},
	}
	di.SubscriberHander = &handlers.SubscriberHandler{
		CampaignRepository:   &database.Repository[campaigns.Campaign]{DB: db},
		SubscriberRepository: &database.Repository[campaigns.Subscriber]{DB: db},
		ClientRepository:     &database.Repository[clients.Client]{DB: db},
	}
	di.ClientHandler = &handlers.ClientHandler{
		UserRepository:   &database.Repository[users.User]{DB: db},
		ClientRepository: &database.Repository[clients.Client]{DB: db},
	}
	return di
}
