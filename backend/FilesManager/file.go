package filesmanager

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type convector struct {
	List FileInfo
}

var OsTree = Dir{}

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

		info := FileInfo{
			Name: file.Name(),
			Dir:  dir,
		}

		list = append(list, info)
	}
	return list, nil
}

func BytesFile(path string) (*ByteFile, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &ByteFile{Name: path, Content: file}, nil
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
