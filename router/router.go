package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"

	"sample/controller"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", controller.CreateUser)
		api.GET("/users", controller.ListUsers)
	}

	// Example usage of Gorilla Mux for something like websockets or advanced routing
	r.GET("/mux", gin.WrapH(muxRouter()))

	return r
}

func muxRouter() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc("/mux", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Gorilla Mux!"))
	}).Methods("GET")
	return m
}
