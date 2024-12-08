package filesmanager

import (
	"FilesWithTag/pkg/path"
	"FilesWithTag/pkg/set"
	"fmt"
	"os"
	pathGo "path"
	"strings"
	"sync"
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

func FilesInDir(dir string, count, offset int, ftype string) ([]Content, []string, error) {
	satrtT := time.Now()
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}
	var list []Content
	tagsDir := set.NewSet()
	wg := sync.WaitGroup{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if ftype != "All" {
			if !strings.Contains(Ftype[ftype], pathGo.Ext(file.Name())) {
				continue
			}
		}

		if offset > 0 {
			offset -= 1
			continue
		}

		pathToFile := strings.Replace(dir+"/"+file.Name(), "\\", "/", -1)

		wg.Add(1)
		go asyncFilesInDir(pathToFile, &list, tagsDir, &wg, &sync.Mutex{})

		count -= 1
		if count == 0 {
			break
		}
	}
	wg.Wait()
	fmt.Println("[DEBUG] ", time.Since(satrtT).Seconds(), "sec")
	return list, tagsDir.ToSlice(), nil
}

// NOTE: кажется логика дырявая, но пока не ясно где и как улучшить
func SearchFileWithTags(dir string, tagsList []string) ([]Content, error) {
	files := make([]string, 0)

	p, err := path.ParsePath(dir)
	if err != nil {
		return nil, err
	}

	dir = p.StringLinux()

	for _, v := range tagsList {
		for _, file := range tags.Get(v) {
			if strings.Contains(file, dir) {
				fmt.Println("[DEBUG]", strings.TrimPrefix(file, dir), file, dir)
				fmt.Println("[DEBUG]", len(strings.Split(strings.TrimPrefix(file, dir), "/")))
				if len(strings.Split(strings.TrimPrefix(strings.TrimPrefix(file, dir), "/"), "/")) == 1 {
					fmt.Println("[DEBUG] finded")
					files = append(files, file)
				}
			}
		}
	}

	res := asyncLoadFilesList(files)

	return res, nil
}

func SearchAllFileWithTags(tagsList []string) ([]Content, error) {
	fmt.Println("[DEBUG] SearchAllFileWithTags", tagsList)
	files := make([]string, 0)

	for _, v := range tagsList {
		files = append(files, tags.Get(v)...)
	}

	res := asyncLoadFilesList(files)

	return res, nil
}
