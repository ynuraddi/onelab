package controller

func (s *Server) routes() {
	s.App.POST("/user", s.handler.CreateUser)
	s.App.GET("/user/:id", s.handler.GetUser)
	s.App.PATCH("/user/:id", s.handler.UpdateUser)
	s.App.DELETE("/user/:id", s.handler.DeleteUser)
}
