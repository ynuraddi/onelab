package postgre

import (
	"log"

	"app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=onelab password=onelab dbname=onelab host=localhost port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
