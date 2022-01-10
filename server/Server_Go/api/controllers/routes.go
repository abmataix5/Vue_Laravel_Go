package controllers

import (
	"github.com/victorsteven/forum/api/middlewares"
)

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api")
	{
		v1.GET("/workers/", s.GetUsers)
		v1.GET("/workers/:id", s.GetUser)
		v1.POST("/workers/", s.CreateUser)
		v1.POST("/login", s.Login)
		v1.PUT("/workers/:id", middlewares.TokenAuthMiddleware(), middlewares.AdminAuthMiddleware(), s.UpdateUser)    /* Comprueba el token y si es admin */
		v1.DELETE("/workers/:id", middlewares.TokenAuthMiddleware(), middlewares.AdminAuthMiddleware(), s.DeleteUser) /* Comprueba el token y si es admin */

	}
}
