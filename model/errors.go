package model

import "errors"

var (
	ErrInvalidJSON         = errors.New("invalid json")
	ErrInvalidData         = errors.New("invalid data")
	ErrInternalServerError = errors.New("unexpected error")

	ErrUserIsNotExist     = errors.New("user is not exist")
	ErrUserIsAlreadyExist = errors.New("user is already exist")
	ErrUserIsNotVerified  = errors.New("user is not verified")

	ErrBookIsNotExist     = errors.New("book is not exist")
	ErrBookIsAlreadyExist = errors.New("book is already eixst")

	ErrBookBorrowIsNotExist = errors.New("record is not exist")

	ErrJWTTokenExpired = errors.New("token is expired")

	ErrContextExceed = errors.New("timeout exceeded")
	ErrEditConflict  = errors.New("edit conflict")
)
