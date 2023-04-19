package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// ListLibraryDebtor godoc
// @Summary List library debtors
// @Description	List library debtor record
// @Tags library
// @Accept json
// @Produce json
// @Success 302 {object} []model.LibraryDebtor
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /library/debtor/list [get]
func (h *Manager) ListBookBorrowDebtor(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.service.Library.ListDebtors(ctx)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, list)
}

// ListLibraryMetric godoc
// @Summary List library metric
// @Description	List library debtor record
// @Tags library
// @Accept json
// @Produce json
// @Param id path int true "Month"
// @Success 200 {object} []model.LibraryMetric
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /library/metric/list/{id} [get]
func (h *Manager) ListBookBorrowMetric(c echo.Context) error {
	input := struct {
		Month int `param:"id" validate:"required,min=1,max=12"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.service.Library.ListMetric(ctx, input.Month)
	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, list)
}
