package handler

import (
	"context"
	"net/http"
	"time"

	"app/model"

	"github.com/labstack/echo/v4"
)

// TransactionPay godoc
// @Summary TransationPay
// @Description	You can pay your order here
// @Tags transaction
// @Accept json
// @Produce json
// @Param input body model.PayTransactionRq true "your transaction and sum amount"
// @Success 200 {object} MsgEnvelope "success transfer"
// @Failure 400 {object} ErrEnvelope "bad request"
// @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// @Failure 500 {object} ErrEnvelope "internal server error"
// @Router /transaction [post]
func (h *Manager) PayTransaction(c echo.Context) error {
	var input model.PayTransactionRq

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	err := h.service.Trans.Pay(ctx, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrEnvelope{err.Error()})
	}

	return c.JSON(http.StatusFound, MsgEnvelope{"success transfer"})
}
