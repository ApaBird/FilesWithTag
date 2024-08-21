package main

import (
	"FilesWithTag/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := Readconfig("config.json")

	r := mux.NewRouter()

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.PathPrefix("/view").HandlerFunc(service.ViewHandler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, r)
}
