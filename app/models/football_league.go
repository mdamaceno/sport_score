package models

import (
	"time"

	"github.com/google/uuid"
)

type FootballLeague struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null;unique" json:"slug"`
	CountryId uuid.UUID `gorm:"type:uuid;not null;" json:"country_id"`
	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;" json:"updated_at"`

	Country Country
}
