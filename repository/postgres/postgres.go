package postgres

import (
	"fmt"
	"log"

	"app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			conf.Database.User, conf.Database.Pass, conf.Database.DBName, conf.Database.Host, conf.Database.Port),
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
