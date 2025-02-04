package routes

import (
	"github.com/chiltom/pogo_buddy/handlers"
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/session-check", handlers.SessionCheckHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
}
