package config

import (
	"github.com/spf13/viper"
)

type AppConf struct {
	APP_ID  string
	DB_HOST string
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_PORT string
}

func InitConfig() *AppConf {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return &AppConf{
		APP_ID:  viper.GetString("APP_ID"),
		DB_HOST: viper.GetString("DB_HOST"),
		DB_USER: viper.GetString("DB_USER"),
		DB_PASS: viper.GetString("DB_PASS"),
		DB_NAME: viper.GetString("DB_NAME"),
		DB_PORT: viper.GetString("DB_PORT"),
	}
}
