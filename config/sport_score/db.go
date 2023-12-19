package config

import (
	"log"

	"github.com/mdmaceno/sport_score/app/sport_score/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := GetEnvVar("DB_HOST").(string)
	user := GetEnvVar("DB_USER").(string)
	password := GetEnvVar("DB_PASS").(string)
	dbname := GetEnvVar("DB_NAME").(string)
	port := GetEnvVar("DB_PORT").(string)

	url := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port +
		" sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Country{})

	return db
}
