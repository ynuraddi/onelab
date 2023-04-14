package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/internal/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Login godoc
//
//	@Summary		Login
//	@Tags			auth
//	@Description	login
//	@Accept			json
//	@Produce		json
//	@Param			input	body	model.LoginUserRq	true	"user login request"
//	@Success		200		{json}	handler.envelope
//	@Failure		400		{json}	handler.envelope
//	@Failure		404		{json}	handler.envelope
//	@Failure		500		{json}	handler.envelope
//	@Router			/login [post]
func (h *Manager) LoginUser(c echo.Context) error {
	var input model.LoginUserRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(CreateUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{
			Msg: "handler(LoginUser): validation failed " + err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	if err := h.s.User.Auth(ctx, model.User{
		Name:     input.Name,
		Password: input.Password,
	}); err != nil {
		return c.JSON(http.StatusUnauthorized, Envelope{Msg: "handler(LoginUser): " + err.Error()})
	}

	token, err := h.s.JWT.GenerateToken(input.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(LoginUser): generate token failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *Manager) CreateUser(c echo.Context) error {
	var input model.UserCreateRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(CreateUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{
			Msg: "handler(CreateUser): validation failed " + err.Error(),
		})
	}

	user := model.User{
		Name:     input.Name,
		Login:    input.Login,
		Password: input.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Create(ctx, user)
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return c.JSON(http.StatusUnprocessableEntity, Envelope{Msg: "handler(CreateUser): user already exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(CreateUser): " + err.Error()})
	default:
		return c.JSON(http.StatusCreated, Envelope{Msg: "user created"})
	}
}

func (h *Manager) GetUser(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(GetUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{
			Msg: "handler(GetUser): validation failed " + err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.s.User.Get(ctx, input.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.JSON(http.StatusNotFound, Envelope{Msg: "handler(GetUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(GetUser): " + err.Error()})
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
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(UpdateUser): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{
			Msg: "handler(UpdateUser): validation failed " + err.Error(),
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
		return c.JSON(http.StatusNotFound, Envelope{Msg: "handler(UpdateUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(UpdateUser): " + err.Error()})
	default:
		return c.JSON(http.StatusOK, Envelope{Msg: "user updated"})
	}
}

func (h *Manager) DeleteUser(c echo.Context) error {
	input := struct {
		ID int `param:"id"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(DeleteUser): bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.User.Delete(ctx, input.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrMissingWhereClause):
		return c.JSON(http.StatusNotFound, Envelope{Msg: "handler(DeleteUser): user is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(DeleteUser): " + err.Error()})
	default:
		return c.JSON(http.StatusOK, Envelope{Msg: "user deleted"})
	}
}
