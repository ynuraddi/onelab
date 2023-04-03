package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"app/internal/model"
	"app/internal/service"
)

type createUserHandler struct {
	srv interface {
		Create(data model.User) error
	}
}

func NewCreateUserHandler(srv service.IUserService) http.Handler {
	return &createUserHandler{
		srv: srv,
	}
}

func (h *createUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input := struct {
		FIO      string `json:"fio"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		InternalServerError(w, fmt.Errorf("createUserHandler: %w", err))
		return
	}

	// TODO validation

	dto := model.User{
		FIO:      input.FIO,
		Login:    input.Login,
		Password: input.Password,
	}

	err := h.srv.Create(dto)
	switch {
	case errors.Is(err, model.ErrUserIsAlreadyExists):
		BadRequest(w, fmt.Errorf("createUserHandler: %w", err))
		return
	case err != nil:
		InternalServerError(w, fmt.Errorf("createUserHandler: %w", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created\n"))
}

type getUserHandler struct {
	srv interface {
		Get(id int) (model.User, error)
	}
}

func NewGetUserHandler(srv service.IUserService) http.Handler {
	return &getUserHandler{
		srv: srv,
	}
}

func (h *getUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		BadRequest(w, err)
		return
	}

	user, err := h.srv.Get(int(id))
	switch {
	case errors.Is(err, model.ErrUserNotExists):
		NotFound(w, r)
		return
	case err != nil:
		InternalServerError(w, err)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Write(data)
}
