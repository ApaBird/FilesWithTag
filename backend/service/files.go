package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"fmt"
	"net/http"
)

type FilesInfo struct {
	Files []filesmanager.FileInfo
}

func FilesHandler(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("F")
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	files, err := filesmanager.FilesInDir(path)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	return FilesInfo{Files: files}
}

func GetFileByte(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file, err := filesmanager.BytesFile(path)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	return *file
}
