package service

import (
	filesmanager "FilesWithTag/file_manager"
	"net/http"
)

type ResponceDirs struct {
	Dirs []filesmanager.Dir `json:"dirs"`
}

// @Summary		Папки в папке
// @Tags		dir
// @Description	Папки в папке
// @ID			getDirs
// @Accept		json
// @Produce		json
// @Param		Path	query		string			true	"Путь до папки"
// @Success		200			{object}	ResponceDirs		"tags"
// @Failure		400,500		{object}	ResponceError		"error"
// @Router		/Dirs [get]
func GetDirs(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")

	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	dirs, err := filesmanager.GetDirs(path)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	responce := ResponceDirs{Dirs: dirs}

	return responce
}
