package main

import (
	//filesmanager "FilesWithTag/FilesManager"

	filesmanager "FilesWithTag/FilesManager"
	"FilesWithTag/service"
	"fmt"
	"net/http"

	_ "FilesWithTag/docs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// TODO: Собирать мета теги и хранить

//	@title			FilesWithTag API
//	@version		2.0
//	@description	FilesWithTag API
//	@host			localhost:8050
//	@BasePath		/

func main() {
	config := Readconfig("config.json")

	filesmanager.AnalyzeStorage()

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"}, // Allowing only get, just an example
	})

	r.Use(InfoRequest)
	SwaggerRouting(r)

	r.HandleFunc("/Files", service.Wrapper(service.FilesHandler)).Methods("GET")
	r.HandleFunc("/FileByte", service.Wrapper(service.GetFileByte)).Methods("GET")
	r.HandleFunc("/OsTree", service.Wrapper(service.OsTreeHandler)).Methods("GET")
	r.HandleFunc("/GetMeta", service.Wrapper(service.GetTags)).Methods("GET")
	r.HandleFunc("/AddMeta", service.Wrapper(service.AddTags)).Methods("POST")
	r.PathPrefix("/").HandlerFunc(service.ViewHandler)

	fmt.Println("Сервер запущен")
	http.ListenAndServe(":"+config.Port, c.Handler(r))
}

func InfoRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func SwaggerRouting(router *mux.Router) {
	prefix := "/docs"
	router.PathPrefix(prefix).Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		//httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		//httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

}
