package app

import (
	"P2/app/controllers"

	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoute() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

}
