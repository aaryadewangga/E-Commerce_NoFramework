package app

import "P2/app/controllers"

func (server *Server) InitializeRoute() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

}
