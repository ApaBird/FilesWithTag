package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"fmt"
	"net/http"
	"strconv"
)

type FilesInfo struct {
	Files []filesmanager.Content
}

func FilesHandler(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("F")
	path := r.URL.Query().Get("Path")
	count := r.URL.Query().Get("Count")
	offset := r.URL.Query().Get("Offset")
	if path == "" || count == "" || offset == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		return ResponceError{Error: ErrNotCorrectTypeParametr.Error(), Status: http.StatusBadRequest}
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return ResponceError{Error: ErrNotCorrectTypeParametr.Error(), Status: http.StatusBadRequest}
	}

	files, err := filesmanager.FilesInDir(path, countInt, offsetInt)
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

	file, b, err := filesmanager.OpenFile(path)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	responce := struct {
		FileName string
		Content  []byte
	}{
		FileName: file.Name,
		Content:  b,
	}

	return responce
}
