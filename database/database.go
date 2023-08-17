package database

import (
	"fmt"
	"log"

	"github.com/bruandreo/order-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(config config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		config.DB_Host,
		config.DB_User,
		config.DB_Password,
		config.DB_Name,
		config.DB_Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect with database", err)
	}

	return db
}
