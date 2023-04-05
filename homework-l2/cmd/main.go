package main

import (
	"log"

	"app/config"
	"app/internal/controller"
	"app/internal/controller/handler"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/client/postgre"
	"app/pkg/validator"
)

func main() {
	config := config.GetConfing()

	// logger := logger.NewLogger(config)

	db := postgre.OpenDB(config)

	repo := repository.NewRepository(db)

	serv := service.NewService(repo)

	validator := validator.NewValidator()

	handlers := handler.NewManager(config, serv)

	server := controller.NewServer(config, handlers, validator)

	log.Fatalln(server.Serve())
}
