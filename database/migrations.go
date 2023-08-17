package database

import (
	"fmt"
	"log"

	"github.com/bruandreo/order-service/internal/config"
	"github.com/tanimutomo/sqlfile"
)

func RunMigrations(config config.Config) {
	db := Connection(config)

	basicDB, err := db.DB()
	if err != nil {
		log.Fatal("Fail to get basic sql db", err)
	}

	sql := sqlfile.New()

	err = sql.Directory("./database/migrations")
	if err != nil {
		log.Fatal("Fail to get migrations directory", err)
	}

	_, err = sql.Exec(basicDB)
	if err != nil {
		log.Fatal("Fail to execute SQL", err)
	}

	fmt.Println("Migrations executed with success")
}
