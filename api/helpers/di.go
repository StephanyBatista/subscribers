package helpers

import (
	"subscribers/infra/database"
	"subscribers/web/handlers"

	"gorm.io/gorm"
)

type DI struct {
	DB                 *gorm.DB
	TokenHandler       *handlers.TokenHandler
	UserHandler        *handlers.UserHandler
	HealthCheckHandler *handlers.HealthCheckHandler
}

func NewDI() *DI {
	di := &DI{}
	db := database.CreateConnection()
	di.DB = db
	di.TokenHandler = &handlers.TokenHandler{Db: db}
	di.UserHandler = &handlers.UserHandler{Db: db}
	di.HealthCheckHandler = &handlers.HealthCheckHandler{Db: db}
	return di
}

func NewFakeDI() *DI {
	FakeEnvs()
	return NewDI()
}
