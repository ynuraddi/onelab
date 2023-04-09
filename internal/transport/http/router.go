package transport

func (s *Server) setupRoutes() {
	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUser)
	s.App.PATCH("/user/:id", s.handler.UpdateUser)
	s.App.DELETE("/user/:id", s.handler.DeleteUser)

	s.App.POST("/books", nil)
	s.App.GET("/book:id", nil)

	s.App.POST("/borrow-history", nil)
	s.App.GET("/borrow-history/:id", nil)
	s.App.GET("/borrow-history/debtors", nil)
	s.App.GET("/borrow-history/stat-month", nil)
}
