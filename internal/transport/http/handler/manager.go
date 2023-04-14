package handler

import (
	"app/config"
	"app/internal/service"
)

type Manager struct {
	conf *config.Config
	s    *service.Service
}

type Envelope struct {
	Msg string `json:"message"`
}

func NewManager(conf *config.Config, service *service.Service) *Manager {
	return &Manager{
		conf: conf,
		s:    service,
	}
}
