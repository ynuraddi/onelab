package transport

import (
	_ "app/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) setupRoutes() {
	s.App.GET("/swagger/*", echoSwagger.WrapHandler)

	// s.App.Use(middleware.LoggerMiddleware)
	auth := s.handler.Auth.ValidateAuth

	s.App.POST("/login", s.handler.LoginUser)

	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUser, auth)
	s.App.PATCH("/user/:id", s.handler.UpdateUser, auth)
	s.App.DELETE("/user/:id", s.handler.DeleteUser, auth)

	s.App.POST("/book", s.handler.CreateBook)
	s.App.GET("/book/:id", s.handler.GetBook)
	s.App.PATCH("/book/:id", s.handler.UpdateBook)
	s.App.DELETE("/book/:id", s.handler.DeleteBook)

	s.App.POST("/book/borrow", s.handler.CreateBookBorrow)
	s.App.GET("/book/borrow/:id", s.handler.GetBookBorrow)
	s.App.PATCH("/book/borrow/:id", s.handler.UpdateBookBorrow)
	s.App.DELETE("/book/borrow/:id", s.handler.DeleteBookBorrow)

	s.App.GET("/book/borrow/debtor/list", s.handler.ListBookBorrowDebtor)
	s.App.GET("/book/borrow/metric/list/:id", s.handler.ListBookBorrowMetric)
}
