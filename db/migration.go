package db

import (
	"github.com/mdmaceno/sport_score/app/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Country{}, &models.FootballLeague{})
}
