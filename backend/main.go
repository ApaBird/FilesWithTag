package main

import (
	"PhotoWithTagServer/service"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", service.HomeHandler).Methods("GET")
}
