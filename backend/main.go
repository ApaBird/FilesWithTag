package main

import (
	filesmanager "FilesWithTag/FilesManager"
	"FilesWithTag/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := Readconfig("config.json")

	filesmanager.AnalyzeStorage()

	r := mux.NewRouter()

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/OsTree", service.Wrapper(service.OsTreeHandler)).Methods("GET")
	r.PathPrefix("/view").HandlerFunc(service.ViewHandler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, r)
}
