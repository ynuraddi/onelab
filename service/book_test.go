package service

import (
	"context"
	"testing"

	"app/model"
	mocks "app/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBookService(t *testing.T) {
	bookRepository := mocks.NewIBookRepository(t)

	bookService := NewBookService(bookRepository)

	cases := []struct {
		name         string
		input        model.CreateBookRq
		want         error
		mockRepoFunc interface{}
	}{
		{
			name: "Create",
			input: model.CreateBookRq{
				Title:  "aboba",
				Author: "aboba",
			},
			want: nil,
			mockRepoFunc: func(ctx context.Context, b model.CreateBookRq) error {
				return nil
			},
		},
		{
			name: "Duplicate",
			input: model.CreateBookRq{
				Title:  "aboba",
				Author: "aboba",
			},
			want: model.ErrBookIsAlreadyExist,
			mockRepoFunc: func(ctx context.Context, b model.CreateBookRq) error {
				return model.ErrBookIsAlreadyExist
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			bookRepository.On("Create", mock.Anything, mock.AnythingOfType("model.CreateBookRq")).
				Return(test.mockRepoFunc).Once()

			get := bookService.Create(context.Background(), test.input)

			assert.ErrorIs(t, get, test.want)
		})
	}
}
