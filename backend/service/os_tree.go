package service

import (
	filesmanager "FilesWithTag/FilesManager"
	"net/http"
)

func OsTreeHandler(w http.ResponseWriter, r *http.Request) any {
	return filesmanager.OsTree
}
