package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"app/config"
	"app/repository"
	"app/service"
	transport "app/transport/http"
	"app/transport/http/handler"
	"app/validator"
)

//	@title			Onelab API - library
//	@version		1.0
//	@description	Onelab project - library

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	log.Fatalf("Service shutdown: %s\n", run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)

	config := config.GetConfing()

	repo := repository.NewRepository(config)

	serv := service.NewService(config, repo)

	validator := validator.NewValidator()

	handlers := handler.NewManager(config, serv)

	HTTPServer := transport.NewServer(config, handlers, validator)

	return HTTPServer.Serve(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)

	go func() {
		log.Println(<-osC)
		c()
	}()
}
