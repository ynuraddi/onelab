package handler

import (
	"context"
	"net/http"
	"time"

	"app/internal/model"

	"github.com/labstack/echo/v4"
)

func (h *Manager) CreateBorrowHistory(c echo.Context) error {
	input := struct {
		BookID     int       `json:"book_id"     validate:"required"`
		UserID     int       `json:"user_id"     validate:"required"`
		BorrowDate time.Time `json:"borrow_date" validate:"required"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(AddEntryToBorrowHistory): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{
			"error":         "handler(AddEntryToBorrowHistory): validation failed",
			"errorValidate": err.Error(),
		})
	}

	entry := model.BookBorrowHistory{
		BookID:     input.BookID,
		UserID:     input.UserID,
		BorrowDate: input.BorrowDate,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.BookBorrow.Create(ctx, entry)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(AddEntryToBorrowHistory): " + err.Error()})
	default:
		return c.JSON(http.StatusCreated, envelope{"info": "entry created"})
	}
}

func (h *Manager) ListDebtorsBorrowHistory(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	list, err := h.s.BookBorrow.ListDebtors(ctx)
	switch {
	// case errors.Is(err, gorm.ErrRecordNotFound):
	// 	return c.JSON(http.StatusNotFound, envelope{"info": "no record"})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(ListDebtorsBorrowHistory)" + err.Error()})
	default:
		return c.JSON(http.StatusFound, list)
	}
}

func (h *Manager) StatMonthBorrowHistory(c echo.Context) error {
	input := struct {
		Month int `json:"month" validate:"required,min=1,max=12"`
		// Year  int `json:"year"  validate:"required,min=2000"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{"error": "handler(StatMonthBorrowHistory): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, envelope{
			"error":         "handler(StatMonthBorrowHistory): validation failed",
			"errorValidate": err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	list, err := h.s.BookBorrow.BookRentalForMonth(ctx, input.Month, 0)
	switch {
	// case errors.Is(err, gorm.ErrRecordNotFound):
	// 	return c.JSON(http.StatusNotFound, envelope{"info": "no record"})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, envelope{"error": "handler(StatMonthBorrowHistory)" + err.Error()})
	default:
		return c.JSON(http.StatusFound, list)
	}
}
