package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user", nil)
	router.HandlerFunc(http.MethodGet, "/user/:id", nil)
	router.HandlerFunc(http.MethodPatch, "/user/:id", nil)
	router.HandlerFunc(http.MethodDelete, "/user/:id", nil)

	return router
}

type Controller struct{}
