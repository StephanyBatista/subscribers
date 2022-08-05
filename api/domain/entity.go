package domain

import (
	"time"

	"github.com/rs/xid"
)

type Entity struct {
	ID        string    `json:"id" gorm:"primaryKey;size:25;"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
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
