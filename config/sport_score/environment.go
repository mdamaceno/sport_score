package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvVar(key string) any {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return viper.Get(key)
}
