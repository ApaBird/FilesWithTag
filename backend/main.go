package main

import (
	"PhotoWithTagServer/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/view/{source}", service.ViewHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
