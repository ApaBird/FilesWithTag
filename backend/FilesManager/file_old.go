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

type Content struct {
	Name    string
	Tags    []string
	Content []byte
}

var OsTree = Dir{}

func FilesInDir(dir string, count, offset int) ([]Content, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []Content
	for _, file := range files {
		if file.IsDir() {
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
			Tags:    f.GetTags(),
			Content: b,
		}

		list = append(list, res)
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
