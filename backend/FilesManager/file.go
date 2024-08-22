package filesmanager

import (
	"fmt"
	"os"
	"strings"

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

func AnalyzeStorage() (ans map[string]*interface{}) {

	ans = make(map[string]*interface{})
	m := make(map[string]*interface{})
	ans["C:/"] = &m
	pull := make([]string, 0)
	pull = append(pull, "C:/")

	for len(pull) > 0 {
		dir := pull[0]
		pull := pull[1:]
		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				split := strings.Split(dir, "/")[1:]
				var target interface{}
				target = ans["C:/"]
				for _, v := range split {
					target = target.(map[string]*interface{})[v]
				}
				*target.(map[string]*interface{})[file.Name()] = make(map[string]*interface{})
				pull = append(pull, dir+"/"+file.Name())
				continue
			}
		}
	}

	return ans
}
