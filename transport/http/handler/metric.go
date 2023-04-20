package handler

import (
	"context"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// ListLibraryMetric godoc
// @Summary ListBookBorrowMetric
// @Description	List library debtor record
// @Tags metric
// @Accept json
// @Produce json
// @Param id path int true "Month"
// @Success 200 {object} []model.LibraryMetricUserBook
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /library/metric/book-per-month/{id} [get]
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

	list, err := h.service.Library.ListMetricBorrow(ctx, input.Month)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, list)
}

// MetricBookTransaction
// @Summary MetricBookTransaction
// @Description	A list of books that clients have now and the total income from each
// @Tags metric
// @Accept json
// @Produce json
// @Success 200 {object} []model.MetricTransactionRp
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 404 {object} ErrEnvelope "record is not exist"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /library/metric/book-amount [get]
func (h *Manager) MetricBookTransaction(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	list, err := h.service.Library.ListMetricTransaction(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, list)
}
