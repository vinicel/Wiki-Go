package server

import (
	"github.com/Wiki-Go/controllers"
	"github.com/gorilla/mux"
)

func (s *Server) InitialiseRoutes(controller *controllers.Controller) {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", controllers.HelloWorldHandler).Methods("GET")
	s.Router.HandleFunc("/articles", controller.GetAllArticles).Methods("GET")
}