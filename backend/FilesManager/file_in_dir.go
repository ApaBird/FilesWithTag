package filesmanager

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
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

func AnalyzeStorage(startDir string) {
	t := time.Now()
	pull := make([]string, 0)

	startDir = strings.Replace(startDir, "\\", "/", -1)
	startDir = strings.Trim(startDir, "/")
	pull = append(pull, startDir)
	OsTree = NewDir(startDir, startDir)

	fmt.Println("[DEBUG]", startDir)

	for len(pull) > 0 {
		dir := pull[0]
		pull = pull[1:]
		if strings.Contains(dir, "Windows") || strings.Contains(dir, "Program Files") || strings.Contains(dir, "ProgramData") {
			continue
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("[ERROR]", err.Error())
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				pull = append(pull, strings.Trim(dir, "/")+"/"+file.Name())
				fmt.Println("[DEBUG]", strings.Trim(dir, "/")+"/"+file.Name())
				d := OsTree.FindDir(dir)
				// fmt.Println("=>", strings.Count(dir, "/"))
				// if strings.Count(dir, "/") >= 3 {
				// 	time.Sleep(time.Second * 10)
				// }
				if d == nil {
					fmt.Println("[ERROR] dir not found")
					continue
				}
				d.AddDirByName(file.Name())
			}
		}
	}

	fmt.Println("Времения потрачено:", time.Since(t))
}
