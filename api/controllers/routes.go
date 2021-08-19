package controllers

import "github.com/RaphaelNagato/goresume-api/api/middlewares"

func (s *Server) initializeRoutes() {

	//Users routes
	s.Router.HandleFunc("/messages", middlewares.SetMiddlewareJSON(s.CreateMessage)).Methods("POST")
	s.Router.HandleFunc("/messages", middlewares.SetMiddlewareJSON(s.GetMessage)).Methods("GET")
}
