package handler

import (
	"app/config"
	"app/service"
)

type Manager struct {
	s *service.Manager
}

type MsgEnvelope struct {
	Msg string `json:"message"`
}

type ErrEnvelope struct {
	Err string `json:"error"`
}

func NewManager(conf *config.Config, service *service.Manager) *Manager {
	return &Manager{
		s: service,
	}
}
