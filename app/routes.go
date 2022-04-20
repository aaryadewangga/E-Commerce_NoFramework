package app

import (
	"P2/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) InitializeRoute() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFileDirectory := http.Dir("app/assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
