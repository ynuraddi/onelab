package postgres

import (
	"context"
	"testing"

	"app/model"

	"github.com/stretchr/testify/assert"
)

func TestCreateBookRepository(t *testing.T) {
	bookRepository := NewBookRepository(testDB)
	ctx := context.Background()

	cases := []struct {
		name  string
		input model.CreateBookRq
		want  error
	}{
		{
			name: "Created",
			input: model.CreateBookRq{
				Title:  "DROP TABLE books;",
				Author: "aboba",
			},
			want: nil,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			get := bookRepository.Create(ctx, test.input)
			assert.ErrorIs(t, get, test.want)

			// b, _ := bookRepository.Get(ctx, i)

			// assert.Equal(t, b, test.input)
		})
	}
}
