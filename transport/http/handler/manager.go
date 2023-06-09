package handler

import (
	"app/config"
	"app/service"
	"app/transport/http/middleware"
)

type Manager struct {
	service *service.Manager
	Auth    *middleware.JWTAuth
}

type MsgEnvelope struct {
	Msg string `json:"message"`
}

type ErrEnvelope struct {
	Err string `json:"error"`
}

func NewManager(conf *config.Config, service *service.Manager) *Manager {
	return &Manager{
		service: service,
		Auth:    middleware.NewJWTAuth(conf, service.User),
	}
}
