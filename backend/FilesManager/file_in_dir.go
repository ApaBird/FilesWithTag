package filesmanager

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type convector struct {
	List File
}

type Content struct {
	Name    string
	Ftype   string
	Tags    []string
	Content []byte
}

var OsTree = Dir{}

func FilesInDir(dir string, count, offset int, ftype string) ([]Content, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []Content
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ftype != "All" {
			if !strings.Contains(Ftype[ftype], path.Ext(file.Name())) {
				continue
			}
		}

		if offset > 0 {
			offset -= 1
			continue
		}

		pathToFile := strings.Replace(dir+"/"+file.Name(), "\\", "/", -1)
		f, b, err := OpenFile(pathToFile)
		if err != nil {
			fmt.Println("[ERROR] ", err.Error())
			continue
		}

		res := Content{
			Name:    file.Name(),
			Ftype:   f.ftype,
			Tags:    f.GetTags().ToSlice(),
			Content: b,
		}

		list = append(list, res)
		count -= 1
		if count == 0 {
			break
		}
	}
	return list, nil
}
