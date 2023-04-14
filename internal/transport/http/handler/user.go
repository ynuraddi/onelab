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

// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param input body model.LoginUserRq true "User login input"
// @Success 200 {object} Envelope
// @Failure 400 {object} Envelope
// @Failure 401 {object} Envelope
// @Failure 500 {object} Envelope
// @Router /login [post]
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags user
// @Accept json
// @Produce json
// @Param input body model.UserCreateRq true "User information"
// @Success 201 {object}	Envelope "user created"
// @Failure 400 {object}	Envelope "bad request"
// @Failure 422 {object}	Envelope "user already exist"
// @Router /user [post]
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

// GetUser godoc
// @Summary Get user by id
// @Description	Get user by id in query param
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 302 {object} model.User
// @Failure 400 {object} Envelope "bad request"
// @Failure 401 {object} Envelope "missing or malformed jwt"
// @Failure 404 {object} Envelope "user is not exist"
// @Failure 500 {object} Envelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [get]
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

// UpdateUser godoc
// @Summary Update user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body model.UserUpdateRq false "User information"
// @Success 200 {object} Envelope "user updated"
// @Failure 400 {object} Envelope "bad request"
// @Failure 401 {object} Envelope "missing or malformed jwt"
// @Failure 404 {object} Envelope "user is not exist"
// @Failure 500 {object} Envelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [patch]
func (h *Manager) UpdateUser(c echo.Context) error {
	var user model.UserUpdateRq

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(UpdateUser): bad request"})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{
			Msg: "handler(UpdateUser): validation failed " + err.Error(),
		})
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

// DeleteUser godoc
// @Summary Delete user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} Envelope "user deleted"
// @Failure 400 {object} Envelope "bad request"
// @Failure 401 {object} Envelope "missing or malformed jwt"
// @Failure 404 {object} Envelope "user is not exist"
// @Failure 500 {object} Envelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [delete]
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
