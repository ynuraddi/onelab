package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags book
// @Accept json
// @Produce json
// @Param input body model.CreateBookRq true "Book information"
// @Success 201 {object} MsgEnvelope "Book created"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 422 {object} ErrEnvelope "book already exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /book [post]
func (h *Manager) CreateBook(c echo.Context) error {
	var input model.CreateBookRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.s.Book.Create(ctx, input)
	if errors.Is(err, model.ErrBookIsAlreadyExist) {
		return c.JSON(http.StatusUnprocessableEntity, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusCreated, MsgEnvelope{Msg: model.StatusBookCreated})
}

// GetBook godoc
// @Summary Get book by id
// @Description	Get book by id in query param
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 302 {object} model.Book
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "book is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/{id} [get]
func (h *Manager) GetBook(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	book, err := h.s.Book.Get(ctx, input.ID)
	if errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusFound, book)
}

// UpdateBook godoc
// @Summary Update book by id
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param input body model.UpdateBookRq false "Book information"
// @Success 200 {object} MsgEnvelope "book updated"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "book is not exist"
// @Failure 409 {object} ErrEnvelope "edit conflict"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/{id} [patch]
func (h *Manager) UpdateBook(c echo.Context) error {
	var input model.UpdateBookRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.s.Book.Update(ctx, input)
	if errors.Is(err, model.ErrEditConflict) {
		return c.JSON(http.StatusConflict, ErrEnvelope{Err: err.Error()})
	} else if errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{Msg: model.StatusBookUpdated})
}

// DeleteBook godoc
// @Summary Delete book by id
// @Tags book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} MsgEnvelope "book deleted"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "book is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/{id} [delete]
func (h *Manager) DeleteBook(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.s.Book.Delete(ctx, input.ID)
	if errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{Msg: model.StatusBookDeleted})
}
