package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahandPgr/Email-verifier/middleware"
)

// Initializing routes
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/form", middleware.FormHandler).Methods("POST")
	return r
}
