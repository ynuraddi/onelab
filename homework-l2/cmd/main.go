package main

import (
	"log"

	"app/config"
	"app/internal/controller"
	"app/internal/controller/handler"
	"app/internal/repository"
	"app/internal/service"
	"app/pkg/client/postgre"
)

func main() {
	config := config.GetConfing()

	db := postgre.OpenDB(config)

	repo := repository.NewRepository(db)

	serv := service.NewService(repo)

	handlers := handler.NewManager(config, serv)

	server := controller.NewServer(config, handlers)

	log.Fatalln(server.Serve())
}
