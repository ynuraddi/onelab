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
	input := struct {
		Name     string `json:"name"     validate:"required,min=5"`
		Login    string `json:"login"    validate:"required,min=5"`
		Password string `json:"password" validate:"required,min=5"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(CreateUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{
			"error":         "handler(CreateUser): validation failed",
			"errorValidate": err.Error(),
		})
	}

	user := model.User{
		Name:     input.Name,
		Login:    input.Login,
		Password: input.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Я вот не знаю где и как лучше обрабатывать ошибки и как их лучше стандартизировать, не довело видеть хороших примеров, у меня не правильно мне кажеться сделано, хотелось бы знать как лучше.

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
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(GetUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{
			"error":         "handler(GetUser): validation failed",
			"errorValidate": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.s.User.Get(ctx, input.ID)
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
	input := struct {
		ID    int    `param:"id"      validate:"required,min=1"`
		Name  string `json:"name"     `
		Login string `json:"login"    `
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(UpdateUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{
			"error":         "handler(UpdateUser): validation failed",
			"errorValidate": err.Error(),
		})
	}

	user := model.User{
		ID:    input.ID,
		Name:  input.Name,
		Login: input.Login,
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
	input := struct {
		ID int `param:"id"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(DeleteUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Delete(ctx, input.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrMissingWhereClause):
		return c.JSON(http.StatusNotFound, envelope{"info": "handler(DeleteUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(DeleteUser): " + err.Error()})
	default:
		return c.JSON(http.StatusOK, envelope{"info": "user deleted"})
	}
}
