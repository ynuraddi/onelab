package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// CreateBookBorrow godoc
// @Summary Create a new book borrow
// @Description Create a new book borrow with the input payload
// @Tags book_borrow
// @Accept json
// @Produce json
// @Param input body model.CreateBookBorrowRq true "Book borrow information time:2020-04-17T18:25:43.511Z"
// @Success 201 {object} MsgEnvelope "record created"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /book/borrow [post]
func (h *Manager) CreateBookBorrow(c echo.Context) error {
	var input model.CreateBookBorrowRq

	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.BookBorrow.Create(ctx, input)
	if errors.Is(err, model.ErrUserIsNotExist) || errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusCreated, MsgEnvelope{model.StatusBookBorrowCreated})
}

// GetBookBorrow godoc
// @Summary Get Borrow record by id
// @Description	Get borrow record by id in query param
// @Tags book_borrow
// @Accept json
// @Produce json
// @Param id path int true "BookBorrow ID"
// @Success 302 {object} model.BookBorrow
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /book/borrow/{id} [get]
func (h *Manager) GetBookBorrow(c echo.Context) error {
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

	record, err := h.service.BookBorrow.Get(ctx, input.ID)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusFound, record)
}

// UpdateBookBorrow godoc
// @Summary Update book borrow by id
// @Tags book_borrow
// @Accept json
// @Produce json
// @Param id path int true "BookBorrow ID"
// @Param input body model.UpdateBookBorrowRq false "record update information"
// @Success 200 {object} MsgEnvelope "record updated"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 409 {object} ErrEnvelope "edit conflict"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /book/borrow/{id} [patch]
func (h *Manager) UpdateBookBorrow(c echo.Context) error {
	var input model.UpdateBookBorrowRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.BookBorrow.Update(ctx, input)
	if errors.Is(err, model.ErrEditConflict) {
		return c.JSON(http.StatusConflict, ErrEnvelope{err.Error()})
	} else if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{model.StatusBookBorrowUpdated})
}

// DeleteBookBorrow godoc
// @Summary Delete book borrow record by id
// @Tags book_borrow
// @Accept json
// @Produce json
// @Param id path int true "BookBorrow ID"
// @Success 200 {object} MsgEnvelope "record deleted"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /book/borrow/{id} [delete]
func (h *Manager) DeleteBookBorrow(c echo.Context) error {
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

	err := h.service.BookBorrow.Delete(ctx, input.ID)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{model.StatusBookBorrowDeleted})
}
