package config

import (
	"log"

	"github.com/mdmaceno/sport_score/app/sport_score/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	url := "host=localhost user=postgres password=postgres dbname=sport_score_development port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Country{})

	return db
}
