package models

import (
	"time"

	"github.com/google/uuid"
)

type FootballTeam struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	Slug      string    `gorm:"type:varchar(255);not null;unique"`
	Logo      string    `gorm:"type:varchar(255);"`
	CountryId uuid.UUID `gorm:"type:uuid;not null;"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`

	Country Country
}
