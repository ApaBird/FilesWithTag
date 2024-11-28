package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"encoding/json"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file := filesmanager.OpenFileWithTags(path)

	responce := struct {
		FilName string
		Tags    []string
	}{
		FilName: file.Name,
		Tags:    file.GetTags(),
	}

	return responce
}

func AddTags(w http.ResponseWriter, r *http.Request) any {
	body := struct {
		Path string   `json:"Path"`
		Tags []string `json:"Tags"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusBadRequest}
	}

	if body.Path == "" || len(body.Tags) == 0 {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file := filesmanager.OpenFileWithTags(body.Path)

	for _, tag := range body.Tags {
		file.AddTag(tag)
	}

	return Responce{Status: http.StatusAccepted, Comment: "OK"}
}
