// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "app/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IBookService is an autogenerated mock type for the IBookService type
type IBookService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, book
func (_m *IBookService) Create(ctx context.Context, book model.CreateBookRq) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateBookRq) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IBookService) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *IBookService) Get(ctx context.Context, id int) (model.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (model.Book, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Book); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, book
func (_m *IBookService) Update(ctx context.Context, book model.UpdateBookRq) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateBookRq) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIBookService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBookService creates a new instance of IBookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBookService(t mockConstructorTestingTNewIBookService) *IBookService {
	mock := &IBookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}