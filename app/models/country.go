package models

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
