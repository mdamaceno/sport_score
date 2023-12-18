package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	url := "host=localhost user=postgres password=postgres dbname=sport_score port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
