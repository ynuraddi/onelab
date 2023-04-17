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
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.s.BookBorrow.Create(ctx, input)
	if errors.Is(err, model.ErrUserIsNotExist) || errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusCreated, MsgEnvelope{Msg: model.StatusBookBorrowCreated})
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
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/borrow/{id} [get]
func (h *Manager) GetBookBorrow(c echo.Context) error {
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

	record, err := h.s.BookBorrow.Get(ctx, input.ID)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
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
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/borrow/{id} [patch]
func (h *Manager) UpdateBookBorrow(c echo.Context) error {
	var input model.UpdateBookBorrowRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.s.BookBorrow.Update(ctx, input)
	if errors.Is(err, model.ErrEditConflict) {
		return c.JSON(http.StatusConflict, ErrEnvelope{Err: err.Error()})
	} else if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{Msg: model.StatusBookBorrowUpdated})
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
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/borrow/{id} [delete]
func (h *Manager) DeleteBookBorrow(c echo.Context) error {
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

	err := h.s.BookBorrow.Delete(ctx, input.ID)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusOK, MsgEnvelope{Msg: model.StatusBookBorrowDeleted})
}

// ListBookBorrowDebtor godoc
// @Summary List Borrow debtors
// @Description	List borrow debtor record
// @Tags book_borrow
// @Accept json
// @Produce json
// @Success 302 {object} []model.BookBorrowDebtorRp
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/borrow/debtor/list [get]
func (h *Manager) ListBookBorrowDebtor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.s.BookBorrow.ListDebtors(ctx)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusFound, list)
}

// ListBookBorrowMetric godoc
// @Summary List Borrow metric
// @Description	List borrow debtor record
// @Tags book_borrow
// @Accept json
// @Produce json
// @Param id path int true "Month"
// @Success 200 {object} []model.BookBorrowMetricRp
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /book/borrow/metric/list/{id} [get]
func (h *Manager) ListBookBorrowMetric(c echo.Context) error {
	input := struct {
		Month int `param:"id" validate:"required,min=1,max=12"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{Err: model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.s.BookBorrow.ListMetric(ctx, input.Month)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{Err: err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{Err: model.ErrInternalServerError.Error()})
	}

	return c.JSON(http.StatusFound, list)
}
