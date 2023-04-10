package transport

func (s *Server) setupRoutes() {
	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUser)
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
