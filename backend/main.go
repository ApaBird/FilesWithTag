package main

import (
	filesmanager "FilesWithTag/FilesManager"
	"FilesWithTag/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := Readconfig("config.json")

	a := filesmanager.AnalyzeStorage()
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))

	r := mux.NewRouter()

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.PathPrefix("/view").HandlerFunc(service.ViewHandler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, r)
}
