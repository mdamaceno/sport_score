package config

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mdmaceno/sport_score/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(a *AppConf) *gorm.DB {
	url := "host=" + a.DB_HOST + " user=" + a.DB_USER + " password=" + a.DB_PASS + " dbname=" + a.DB_NAME + " port=" +
		a.DB_PORT + " sslmode=disable TimeZone=UTC"

	orm, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	db.RunMigrations(orm)

	log.Println("Database migrations ran successfully")

	return orm
}

func MockDB() *gorm.DB {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db
}
