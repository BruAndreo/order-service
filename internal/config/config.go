package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	DB_Host     string
	DB_User     string
	DB_Password string
	DB_Port     string
	DB_Name     string
}

func Initialize() (config Config) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error to read env file", err)
	}

	config.AppPort = viper.GetString("PORT")
	config.DB_Host = viper.GetString("POSTGRES_HOST")
	config.DB_Port = viper.GetString("POSTGRES_PORT")
	config.DB_Name = viper.GetString("POSTGRES_DB")
	config.DB_User = viper.GetString("POSTGRES_USER")
	config.DB_Password = viper.GetString("POSTGRES_PASSWORD")
	return
}
