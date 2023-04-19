package postgres

import (
	"fmt"
	"log"
	"os"

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

	migrations, err := os.ReadDir("./repository/postgres/migrations")
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range migrations {
		if file.Name()[len(file.Name())-len("up.sql"):] != "up.sql" {
			continue
		}

		content, err := os.ReadFile("./repository/postgres/migrations/" + file.Name())
		if err != nil {
			log.Fatalln(err)
		}

		if err := db.Exec(string(content)).Error; err != nil {
			log.Fatalln(err)
		}

		log.Printf("success migrate: %s\n", file.Name())
	}

	return db
}
