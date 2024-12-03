package filesmanager

import (
	"os"
	"strings"
)

type Dir struct {
	Name string
	Path string
}

func GetDirs(path string) []Dir {
	var dirs []Dir

	path = strings.Replace(path, "/", "\\", -1)
	files, err := os.ReadDir(path)
	if err != nil {
		return dirs
	}
	for _, f := range files {
		if f.IsDir() {
			p := strings.Trim(path, "\\")
			dirs = append(dirs, Dir{Name: f.Name(), Path: strings.Replace(p+"/"+f.Name(), "\\", "/", -1)})
		}
	}
	return dirs
}
