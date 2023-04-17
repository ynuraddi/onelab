package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"app/config"
	"app/transport/http/handler"

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

func (s *Server) Serve(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.setupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf(s.config.HTTP.Host + ":" + s.config.HTTP.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server Shutdown failed: %s\n", err)
	}

	log.Println("Server gracefully exited")

	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()

	e.Validator = s.validator

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
	}))

	return e
}
