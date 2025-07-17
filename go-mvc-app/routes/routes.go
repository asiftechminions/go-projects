package routes

import (
	"go-mvc-app/go-mvc-app/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
}
