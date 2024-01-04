package tests

import (
	"log"

	"github.com/mdmaceno/sport_score/config"
	"github.com/mdmaceno/sport_score/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(a *config.AppConf) *gorm.DB {
	url := "host=" + a.DB_TEST_HOST + " user=" + a.DB_TEST_USER + " password=" + a.DB_TEST_PASS + " dbname=" +
		a.DB_TEST_NAME + " port=" + a.DB_TEST_PORT + " sslmode=disable TimeZone=UTC"

	orm, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.RunMigrations(orm)

	return orm
}

func TruncateTables(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE countries RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE football_leagues RESTART IDENTITY CASCADE")
	db.Exec("TRUNCATE TABLE football_teams RESTART IDENTITY CASCADE")
}
