package main

import (
	//filesmanager "FilesWithTag/FilesManager"

	pythonmoduleapi "FilesWithTag/PythonModuleAPI"
	"FilesWithTag/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config := Readconfig("config.json")

	//filesmanager.AnalyzeStorage()

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	metaDataModule := pythonmoduleapi.NewModule(config.ModuleIp, config.ModulePort)
	metaDataModule.SendData("GetMeta c:\\Users\\apabi\\Desktop\\ \\Нейронка\\stable-diffusion-webui\\outputs\\works\\sfw\\00005-3381904655.png")
	fmt.Println("=>", metaDataModule.Scan())

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/OsTree", service.Wrapper(service.OsTreeHandler)).Methods("GET")
	r.PathPrefix("/view").HandlerFunc(service.ViewHandler)
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, c.Handler(r))
}
