package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"encoding/json"
	"net/http"
)

type ResponceFile struct {
	FilName string
	Tags    []string
}

type AddTagsRequest struct {
	Path string   `json:"Path"`
	Tags []string `json:"Tags"`
}

// @Summary		Получение тегов
// @Tags			file
// @Description	Получение тегов по пути до файла
// @ID				getTags
// @Accept			json
// @Produce		json
// @Param			Path	query		string			true	"path"
// @Success		200		{object}	ResponceFile	"tags"
// @Failure		400		{object}	ResponceError	"error"
// @Router			/GetMeta [get]
func GetTags(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file := filesmanager.OpenFileWithTags(path)

	responce := ResponceFile{
		FilName: file.Name,
		Tags:    file.GetTags(),
	}

	return responce
}

// @Summary		Добавление тегов
// @Tags			file
// @Description	Добавление тегов по пути до файла
// @ID				addTags
// @Accept			json
// @Produce		json
// @Param			AddTagsRequest	body		AddTagsRequest	true	"path"
// @Success		200		{object}	Responce		"tags"
// @Failure		400		{object}	ResponceError	"error"
// @Router			/AddMeta [post]
func AddTags(w http.ResponseWriter, r *http.Request) any {
	body := AddTagsRequest{}

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
