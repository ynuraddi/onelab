package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"app/config"
	"app/internal/repository"
	"app/internal/service"
	transport "app/internal/transport/http"
	"app/internal/transport/http/handler"
	"app/validator"
)

// TODO настроить линтер
// https://olegk.dev/go-linters-configuration-the-right-version

func main() {
	log.Fatalf("Service shutdown: %s\n", run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)

	config := config.GetConfing()

	// logger := logger.NewLogger(config)

	repo := repository.NewRepository(config)

	serv := service.NewService(repo)

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
