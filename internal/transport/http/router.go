package transport

import (
	"app/internal/model"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	_ "app/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) setupRoutes() {
	s.App.GET("/swagger/*", echoSwagger.WrapHandler)

	auth := echojwt.WithConfig(echojwt.Config{
		ContextKey: "user",
		SigningKey: []byte(s.config.HTTP.JWTKey),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JWTClaims)
		},
	})

	s.App.POST("/login", s.handler.LoginUser)

	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUser, auth)
	s.App.PATCH("/user/:id", s.handler.UpdateUser)
	s.App.DELETE("/user/:id", s.handler.DeleteUser)

	s.App.POST("/book", s.handler.CreateBook)
	s.App.GET("/book/:id", s.handler.GetBook)

	s.App.POST("/borrow-history", s.handler.BorrowBook)
	// s.App.GET("/borrow-history/:id", nil)
	s.App.PATCH("/borrow-history/:id", s.handler.ReturnBook) // returning
	s.App.GET("/borrow-history/debtors", s.handler.ListDebtorsBorrowHistory)
	s.App.GET("/borrow-history/stat-month", s.handler.StatMonthBorrowHistory)
}
