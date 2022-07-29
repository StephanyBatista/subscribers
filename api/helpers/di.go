package helpers

import (
	"subscribers/infra/database"
	"subscribers/web/handlers"

	"gorm.io/gorm"
)

type DI struct {
	DB                 *gorm.DB
	UserHandler        *handlers.UserHandler
	HealthCheckHandler *handlers.HealthCheckHandler
}

func NewDI() *DI {
	di := &DI{}
	db := database.CreateConnection()
	di.DB = db
	di.UserHandler = handlers.NewUserHandler(db)
	di.HealthCheckHandler = handlers.NewHealthCheckHandler(db)
	return di
}

func NewFakeDI() *DI {
	FakeEnvs()
	return NewDI()
}
