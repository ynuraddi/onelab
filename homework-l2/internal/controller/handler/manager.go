package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/config"
	"app/internal/model"
	"app/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (h *Manager) CreateUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(CreateUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Create(ctx, user)
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return c.JSON(http.StatusUnprocessableEntity, envelope{"error": "handler(CreateUser): user already exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(CreateUser): " + err.Error()})
	default:
		return c.JSON(http.StatusCreated, envelope{"info": "user created"})
	}
}

func (h *Manager) GetUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(GetUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.s.User.Get(ctx, user.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.JSON(http.StatusNotFound, envelope{"info": "handler(GetUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(GetUser): " + err.Error()})
	default:
		return c.JSON(http.StatusFound, user)
	}
}

func (h *Manager) UpdateUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(UpdateUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Update(ctx, user)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrMissingWhereClause):
		return c.JSON(http.StatusNotFound, envelope{"info": "handler(UpdateUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(UpdateUser): " + err.Error()})
	default:
		return c.JSON(http.StatusOK, envelope{"info": "user updated"})
	}
}

func (h *Manager) DeleteUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(DeleteUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Update(ctx, user)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrMissingWhereClause):
		return c.JSON(http.StatusNotFound, envelope{"info": "handler(DeleteUser): user is not exist\n" + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(DeleteUser): " + err.Error()})
	default:
		return c.JSON(http.StatusOK, envelope{"info": "user deleted"})
	}
}
