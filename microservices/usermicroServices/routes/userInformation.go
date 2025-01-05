package routes

import (
	"user/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/api/users/adduser", controller.AddUser).Methods("POST")

	return router
}
