package filesmanager

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type convector struct {
	List File
}

var OsTree = Dir{}

func FilesInDir(dir string) ([]File, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []File
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		info := File{
			Name: file.Name(),
			dir:  dir,
		}

		list = append(list, info)
	}
	return list, nil
}

func AnalyzeStorage() {
	t := time.Now()
	pull := make([]string, 0)
	pull = append(pull, "C:/")

	OsTree = NewDir("C:", "C:/")

	for len(pull) > 0 {
		dir := pull[0]
		pull = pull[1:]
		if strings.Contains(dir, "Windows") || strings.Contains(dir, "Program Files") || strings.Contains(dir, "ProgramData") {
			continue
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				pull = append(pull, dir+"/"+file.Name())
				d := OsTree.FindDir(dir)
				// fmt.Println("=>", strings.Count(dir, "/"))
				// if strings.Count(dir, "/") >= 3 {
				// 	time.Sleep(time.Second * 10)
				// }
				if d == nil {
					continue
				}
				d.AddDirByName(file.Name())
			}
		}
	}

	fmt.Println("Времения потрачено:", time.Since(t))
}
