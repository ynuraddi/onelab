package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/config"
	"app/model"
	"app/service"
	"app/transport/http/handler"
	"app/validator"

	mocks "app/service/mock"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBookHandler(t *testing.T) {
	e := echo.New()
	e.Validator = validator.NewValidator()

	bookService := mocks.NewIBookService(t)

	service := service.Manager{
		Book: bookService,
	}

	handler := handler.NewManager(&config.Config{
		JWTKey: "secret",
	}, &service)

	cases := []struct {
		name        string
		input       model.CreateBookRq
		wantMsg     string
		wantStat    int
		needMock    bool
		mockService interface{}
	}{
		{
			name: "Created",
			input: model.CreateBookRq{
				Title:  "aboba",
				Author: "aboba",
			},
			wantMsg:  `{"message":"book created"}`,
			wantStat: http.StatusCreated,
			needMock: true,
			mockService: func(ctx context.Context, u model.CreateBookRq) error {
				return nil
			},
		},
		{
			name: "Duplicate",
			input: model.CreateBookRq{
				Title:  "aboba",
				Author: "aboba",
			},
			wantMsg:  `{"error":"book is already exist"}`,
			wantStat: http.StatusUnprocessableEntity,
			needMock: true,
			mockService: func(ctx context.Context, u model.CreateBookRq) error {
				return model.ErrBookIsAlreadyExist
			},
		},
		{
			name: "Validate",
			input: model.CreateBookRq{
				Title:  "abobab",
				Author: "",
			},
			wantMsg:  `{"error":"invalid data"}`,
			wantStat: http.StatusBadRequest,
			needMock: false,
		},
		{
			name: "Internal",
			input: model.CreateBookRq{
				Title:  "abobab",
				Author: "abobab",
			},
			wantMsg:  `{"error":"unexpected error"}`,
			wantStat: http.StatusInternalServerError,
			needMock: true,
			mockService: func(ctx context.Context, u model.CreateBookRq) error {
				return errors.New("123")
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			jsonBook, _ := json.Marshal(test.input)
			rsp := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(jsonBook))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rsp)

			if test.needMock {
				bookService.On("Create", mock.Anything, mock.AnythingOfType("model.CreateBookRq")).
					Return(test.mockService).Once()
			}

			err := handler.CreateBook(c)
			assert.NoError(t, err)
			assert.JSONEq(t, test.wantMsg, rsp.Body.String())
			assert.Equal(t, test.wantStat, rsp.Code)
		})
	}
}
