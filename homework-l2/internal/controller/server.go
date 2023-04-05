package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app/config"
	"app/internal/controller/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config    *config.Config
	handler   *handler.Manager
	validator echo.Validator
	App       *echo.Echo
}

func NewServer(conf *config.Config, handler *handler.Manager, validator echo.Validator) *Server {
	return &Server{
		config:    conf,
		handler:   handler,
		validator: validator,
	}
}

func (s *Server) Serve() error {
	s.App = echo.New()

	s.App.Validator = s.validator

	s.App.Use(middleware.Logger())
	s.App.Use(middleware.Recover())

	s.App.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
	}))

	s.routes()

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.App.Shutdown(ctx); err != nil {
			shutdownError <- err
			return
		}

		shutdownError <- nil
	}()

	log.Println("Start server: http://" + s.config.HTTP.Host + ":" + s.config.HTTP.Port)

	err := s.App.Start(s.config.HTTP.Host + ":" + s.config.HTTP.Port)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
