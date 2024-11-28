package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file := filesmanager.OpenFile(path)

	responce := struct {
		FilName string
		Tags    []string
	}{
		FilName: file.Name,
		Tags:    file.GetTags(),
	}

	return responce
}
