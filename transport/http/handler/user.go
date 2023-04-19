package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// const handlerUserPath = `handlerUser: %s`

// LoginUser godoc
// @Summary LoginUser
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param input body model.LogInRq true "User login input"
// @Success 200 {object} MsgEnvelope "logined"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "unauthorized"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /login [post]
func (h *Manager) LoginUser(c echo.Context) error {
	var input model.LogInRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.User.Authenticate(ctx, input)
	if errors.Is(err, model.ErrUserIsNotExist) {
		return c.JSON(http.StatusUnauthorized, ErrEnvelope{err.Error()})
	} else if errors.Is(err, model.ErrUserWrongPassword) {
		return c.JSON(http.StatusUnauthorized, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	token, err := h.Auth.GenerateJWT(input.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusCreated, MsgEnvelope{token})
}

// CreateUser godoc
// @Summary CreateUser
// @Description Create a new user with the input payload
// @Tags user
// @Accept json
// @Produce json
// @Param input body model.CreateUserRq true "User information"
// @Success 201 {object} MsgEnvelope "user created"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 422 {object} ErrEnvelope "user already exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /user [post]
func (h *Manager) CreateUser(c echo.Context) error {
	var input model.CreateUserRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.User.Create(ctx, input)
	// WARNING почему то не хэндлиться ошибка в репозитории
	if errors.Is(err, model.ErrUserIsAlreadyExist) {
		return c.JSON(http.StatusUnprocessableEntity, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusCreated, MsgEnvelope{model.StatusUserCreated})
}

// GetUser godoc
// @Summary GetUser
// @Description	Get user by id in query param
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 302 {object} model.User
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "user is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [get]
func (h *Manager) GetUser(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	user, err := h.service.User.Get(ctx, input.ID)
	if errors.Is(err, model.ErrUserIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, user)
}

// UpdateUser godoc
// @Summary UpdateUser
// @Description	Update user by id in query param
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body model.UpdateUserRq false "User information"
// @Success 200 {object} MsgEnvelope "user updated"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "user is not exist"
// @Failure 409 {object} ErrEnvelope "edit conflict"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [patch]
func (h *Manager) UpdateUser(c echo.Context) error {
	var input model.UpdateUserRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.User.Update(ctx, input)
	if errors.Is(err, model.ErrEditConflict) {
		return c.JSON(http.StatusConflict, ErrEnvelope{err.Error()})
	} else if errors.Is(err, model.ErrUserIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{model.StatusUserUpdated})
}

// DeleteUser godoc
// @Summary DeleteUser
// @Description	Delete user by id in query param
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} MsgEnvelope "user deleted"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "user is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /user/{id} [delete]
func (h *Manager) DeleteUser(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.User.Delete(ctx, input.ID)
	if errors.Is(err, model.ErrUserIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{model.StatusUserDeleted})
}
