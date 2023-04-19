package handler

// // ListBookBorrowDebtor godoc
// // @Summary List Borrow debtors
// // @Description	List borrow debtor record
// // @Tags book_borrow
// // @Accept json
// // @Produce json
// // @Success 302 {object} []model.BookBorrowDebtorRp
// // @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// // @Failure 404 {object} ErrEnvelope "record is not exist"
// // @Failure 500 {object} ErrEnvelope "internal server error"
// // @Router /book/borrow/debtor/list [get]
// func (h *Manager) ListBookBorrowDebtor(c echo.Context) error {
// 	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
// 	defer cancel()

// 	list, err := h.service.BookBorrow.ListDebtors(ctx)
// 	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
// 		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
// 	} else if err != nil {
// 		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
// 	}

// 	return c.JSON(http.StatusFound, list)
// }

// // ListBookBorrowMetric godoc
// // @Summary List Borrow metric
// // @Description	List borrow debtor record
// // @Tags book_borrow
// // @Accept json
// // @Produce json
// // @Param id path int true "Month"
// // @Success 200 {object} []model.BookBorrowMetricRp
// // @Failure 400 {object} ErrEnvelope "bad request"
// // @Failure 401 {object} ErrEnvelope "missing or malformed jwt"
// // @Failure 404 {object} ErrEnvelope "record is not exist"
// // @Failure 500 {object} ErrEnvelope "internal server error"
// // @Router /book/borrow/metric/list/{id} [get]
// func (h *Manager) ListBookBorrowMetric(c echo.Context) error {
// 	input := struct {
// 		Month int `param:"id" validate:"required,min=1,max=12"`
// 	}{}

// 	if err := c.Bind(&input); err != nil {
// 		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidJSON.Error()})
// 	}

// 	if err := c.Validate(input); err != nil {
// 		return c.JSON(http.StatusBadRequest, ErrEnvelope{model.ErrInvalidData.Error()})
// 	}

// 	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
// 	defer cancel()

// 	list, err := h.service.BookBorrow.ListMetric(ctx, input.Month)
// 	if errors.Is(err, model.ErrBookBorrowIsNotExist) {
// 		return c.JSON(http.StatusNotFound, ErrEnvelope{err.Error()})
// 	} else if err != nil {
// 		return c.JSON(http.StatusInternalServerError, ErrEnvelope{model.ErrInternalServerError.Error()})
// 	}

// 	return c.JSON(http.StatusFound, list)
// }
