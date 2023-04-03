package controller

import (
	"net/http"

	"app/internal/controller/handler"
	"app/internal/controller/middleware"
	"app/internal/service"

	"github.com/julienschmidt/httprouter"
)

func Routes(service *service.Service) http.Handler {
	router := httprouter.New()

	createUser := handler.NewCreateUserHandler(service.User)
	getUser := handler.NewGetUserHandler(service.User)

	router.HandlerFunc(http.MethodPost, "/user", createUser.ServeHTTP)
	router.HandlerFunc(http.MethodGet, "/user/:id", getUser.ServeHTTP)

	router.NotFound = http.HandlerFunc(handler.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(handler.MethodNotAllowed)

	middleware := middleware.RecoverPanic

	return middleware(router)
}
