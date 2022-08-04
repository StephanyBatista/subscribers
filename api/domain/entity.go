package domain

import (
	"time"

	"github.com/rs/xid"
)

type Entity struct {
	ID        string    `gorm:"primaryKey;size:25;"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (e *Entity) IDIsNull() bool {
	return e == nil || e.ID == ""
}

func NewEntity() Entity {

	return Entity{
		ID:        xid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
