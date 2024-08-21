package main

import (
	"FilesWithTag/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := Readconfig("config.json")

	r := mux.NewRouter()

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/view/{source}", service.ViewHandler).Methods("GET")

	http.ListenAndServe(":"+config.Port, r)
}
