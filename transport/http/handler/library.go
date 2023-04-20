package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// LibraryBorrow godoc
// @Summary LibraryBorrow
// @Description	Create library record to book_borrow
// @Tags library
// @Accept json
// @Produce json
// @Param input body model.LibraryBorrowRq true "Rent book, you can specify the rental time in order to view the metric more conveniently TIME-FROMAT:"2020-04-04""
// @Success 201 {object} MsgEnvelope "record created"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /library [post]
func (h *Manager) LibraryBorrow(c echo.Context) error {
	var input model.LibraryBorrowRq

	login, ok := c.Request().Context().Value(model.ContextLogin).(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, ErrEnvelope{"missing login"})
	}

	input.UserLogin = login

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	invoice, err := h.service.Library.BorrowBook(ctx, input)
	if errors.Is(err, model.ErrUserIsNotExist) || errors.Is(err, model.ErrBookIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, invoice)
}

// LibraryReturn godoc
// @Summary LibraryReturn
// @Description	Return book to library, return day is today
// @Tags library
// @Accept json
// @Produce json
// @Param input body model.LibraryBorrowRq true "Rent book, you can specify the rental time in order to view the metric more conveniently TIME-FROMAT:2020-04-04"
// @Success 201 {object} MsgEnvelope "record updated"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Security ApiKeyAuth
// @In header
// @Name Authorization
// @Router /library [patch]
func (h *Manager) LibraryReturn(c echo.Context) error {
	var input model.LibraryReturnRq

	login, ok := c.Request().Context().Value(model.ContextLogin).(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, ErrEnvelope{"missing login"})
	}

	input.UserLogin = login

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.Library.ReturnBook(ctx, input)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{"you don't owe us this book"})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, MsgEnvelope{"book returned"})
}

// ListLibraryDebtor godoc
// @Summary ListBookBorrowDebtor
// @Description	List library debtor record
// @Tags library
// @Accept json
// @Produce json
// @Success 302 {object} []model.LibraryDebtor
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /library/debtors [get]
func (h *Manager) ListBookBorrowDebtor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.service.Library.ListDebtors(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, list)
}
