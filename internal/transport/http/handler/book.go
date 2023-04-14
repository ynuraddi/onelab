package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"app/internal/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Manager) CreateBook(c echo.Context) error {
	input := struct {
		Title  string `json:"title"     validate:"required"`
		Author string `json:"author"    validate:"required"`
	}{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(CreateBook): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(CreateBook): validation failed " + err.Error()})
	}

	book := model.Book{
		Title:  input.Title,
		Author: input.Author,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.Book.Create(ctx, book)
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return c.JSON(http.StatusUnprocessableEntity, Envelope{Msg: "handler(CreateBook): book already exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(CreateBook): " + err.Error()})
	default:
		return c.JSON(http.StatusCreated, Envelope{Msg: "book created"})
	}
}

func (h *Manager) GetBook(c echo.Context) error {
	input := struct {
		ID int `param:"id" validate:"required,min=1"`
	}{}

	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(GetBook): bad request"})
	}

	if err := c.Validate(input); err != nil {
		return c.JSON(http.StatusBadRequest, Envelope{Msg: "handler(GetBook): validation failed " + err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.s.Book.Get(ctx, input.ID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return c.JSON(http.StatusNotFound, Envelope{Msg: "handler(GetBook): book is not exist || " + err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, Envelope{Msg: "handler(GetBook): " + err.Error()})
	default:
		return c.JSON(http.StatusFound, user)
	}
}
