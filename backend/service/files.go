package service

import (
	filesmanager "FilesWithTag/file_manager"
	"net/http"
	"strconv"
)

type FilesResponce struct {
	TagsInDir []string
	Files     []filesmanager.Content
}

type FileByte struct {
	FileName string
	Content  []byte
}

// @Summary		Получение тегов
// @Tags		file
// @Description	Получение тегов по пути до файла
// @ID			GetFilesInDir
// @Accept		json
// @Produce		json
// @Param		Path	query		string			true	"Путь до папки"
// @Param		Count	query		string			true	"Количество"
// @Param		Offset	query		string			true	"Отступ"
// @Param		Ftype	query		string			false	"Тип файлов" Enums(Image, Music, Video, Text)
// @Success		200		{object}	FilesResponce			"список файлов"
// @Failure		400,500	{object}	ResponceError			"error"
// @Router		/Files [get]
func FilesHandler(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	count := r.URL.Query().Get("Count")
	offset := r.URL.Query().Get("Offset")
	ftype := r.URL.Query().Get("Ftype")
	if path == "" || count == "" || offset == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	if ftype == "" {
		ftype = "All"
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		return ResponceError{Error: ErrNotCorrectTypeParametr.Error(), Status: http.StatusBadRequest}
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return ResponceError{Error: ErrNotCorrectTypeParametr.Error(), Status: http.StatusBadRequest}
	}

	files, tags, err := filesmanager.FilesInDir(path, countInt, offsetInt, ftype)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	return FilesResponce{Files: files, TagsInDir: tags}
}

// @Summary		Получение файла
// @Tags			file
// @Description	Получение файла в формате байт строки
// @ID				GetFilesByte
// @Accept			json
// @Produce		json
// @Param			Path	query		string			true	"Путь до файла"
// @Success		200		{object}	FileByte	"имя файла и его содержимое"
// @Failure		400,500		{object}	ResponceError	"error"
// @Router			/FileByte [get]
func GetFileByte(w http.ResponseWriter, r *http.Request) any {
	path := r.URL.Query().Get("Path")
	if path == "" {
		return ResponceError{Error: ErrParametrs.Error(), Status: http.StatusBadRequest}
	}

	file, b, err := filesmanager.OpenFile(path)
	if err != nil {
		return ResponceError{Error: err.Error(), Status: http.StatusInternalServerError}
	}

	responce := FileByte{
		FileName: file.Name,
		Content:  b,
	}

	return responce
}
