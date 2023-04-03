package model

import "errors"

type User struct {
	ID       int    `json:"id"`
	FIO      string `json:"fio"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

var (
	ErrUserNotExists       = errors.New("User does not exist")
	ErrUserIsAlreadyExists = errors.New("User already exists")
	ErrUserOverflow        = errors.New("User storage is overflow")
)
