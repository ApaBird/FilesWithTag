package filesmanager

import (
	"FilesWithTag/pkg/path"
	"os"
)

type Dir struct {
	Name string
	Path string
}

func GetDirs(dir string) ([]Dir, error) {
	var dirs []Dir

	p, err := path.ParsePath(dir)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(p.String())
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			p.Join(f.Name())
			dirs = append(dirs, Dir{Name: f.Name(), Path: p.StringLinux()})
			p.Back()
		}
	}
	return dirs, nil
}
