package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"net/http"
)

// @Summary		Дерево папок
// @Tags		OsTree
// @Description	Дерево папок, дерево состовляется в момент старта сервера
// @ID			GetOsTree
// @Accept		json
// @Produce		json
// @Success		200	{object}	filesmanager.Dir	"дерево папок"
// @Router		/OsTree [get]
func OsTreeHandler(w http.ResponseWriter, r *http.Request) any {
	return filesmanager.OsTree
}
