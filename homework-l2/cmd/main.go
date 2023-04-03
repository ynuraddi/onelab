package main

import (
	"log"

	"app/internal/controller"
	"app/internal/repository"
	"app/internal/service"
)

func main() {
	repo := repository.NewRepository()

	serv := service.NewService(repo)

	ctrl := controller.Routes(serv)

	log.Println(controller.Serve(ctrl))
}
