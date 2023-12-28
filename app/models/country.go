package models

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	Slug      string    `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}
