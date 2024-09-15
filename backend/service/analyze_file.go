package service

import "net/http"

func GetMetaData(w http.ResponseWriter, r *http.Request) any {
	http.Redirect(w, r, "http://localhost:8051"+r.URL.String(), http.StatusMovedPermanently)

	return nil
}
