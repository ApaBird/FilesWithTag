package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	Dir string
	Tag []string
}

type SearchResponce struct {
	Files []filesmanager.Content
}

func SearchHandler(w http.ResponseWriter, r *http.Request) any {
	request := SearchRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusBadRequest}
	}

	if request.Dir == "" || len(request.Tag) == 0 {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	res, err := filesmanager.SearchFileWithTags(request.Dir, request.Tag)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	responce := SearchResponce{
		Files: res,
	}

	return responce
}
