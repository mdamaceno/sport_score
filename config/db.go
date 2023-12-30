package config

import (
	"log"

	"github.com/mdmaceno/sport_score/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(appConf *AppConf) *gorm.DB {
	host := appConf.DB_HOST
	user := appConf.DB_USER
	password := appConf.DB_PASS
	dbname := appConf.DB_NAME
	port := appConf.DB_PORT

	url := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port +
		" sslmode=disable TimeZone=UTC"

	orm, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	db.RunMigrations(orm)

	log.Println("Database migrations ran successfully")

	return orm
}
