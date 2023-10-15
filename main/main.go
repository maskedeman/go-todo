package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maskedemann/go-todo/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterListRoutes(r)

	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
