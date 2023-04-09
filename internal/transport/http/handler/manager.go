package handler

import (
	"app/config"
	"app/internal/service"
)

type Manager struct {
	conf *config.Config
	s    *service.Service
}

type envelope map[string]interface{}

func NewManager(conf *config.Config, service *service.Service) *Manager {
	return &Manager{
		conf: conf,
		s:    service,
	}
}
