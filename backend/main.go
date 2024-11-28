package main

import (
	//filesmanager "FilesWithTag/FilesManager"

	filesmanager "FilesWithTag/FilesManager"
	"FilesWithTag/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// TODO: Собирать мета теги и хранить

func main() {
	config := Readconfig("config.json")

	filesmanager.AnalyzeStorage()

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/FileByte", service.Wrapper(service.GetFileByte)).Methods("GET")
	r.HandleFunc("/OsTree", service.Wrapper(service.OsTreeHandler)).Methods("GET")
	r.HandleFunc("/GetMeta", service.Wrapper(service.GetTags)).Methods("GET")
	r.HandleFunc("/AddMeta", service.Wrapper(service.AddTags)).Methods("POST")
	r.PathPrefix("/").HandlerFunc(service.ViewHandler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, c.Handler(r))
}
