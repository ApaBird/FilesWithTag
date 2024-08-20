package filesmanager

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type convector struct {
	List FileInfo
}

func (c convector) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Println("Ok")
	c.List.Meta[name] = tag
	return nil
}

func FilesInDir(dir string) ([]FileInfo, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []FileInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		info, err := exif.Decode(f)
		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}
		var c convector
		c.List = FileInfo{
			Name: file.Name(),
			Dir:  dir,
			Meta: make(map[interface{}]interface{}, 0),
		}

		info.Walk(c)

		f.Close()

		list = append(list, c.List)
	}
	return list, nil
}
