package routes

import (
	"github.com/gorilla/mux"
	"github.com/maskedemann/go-todo/pkg/controllers"
)

var RegisterListRoutes = func(router *mux.Router) {
	router.HandleFunc("/create/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/list/", controllers.GetAll).Methods("GET")
	router.HandleFunc("/get/{taskId}/", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/update/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/delete/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
