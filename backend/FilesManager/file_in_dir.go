package filesmanager

import (
	"FilesWithTag/pkg/path"
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

func FilesInDir(dir string, count, offset int, ftype string) ([]Content, error) {
	satrtT := time.Now()
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var list []Content
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
		go asyncFilesInDir(pathToFile, &list, &wg, &sync.Mutex{})

		// f, b, err := OpenFile(pathToFile)
		// if err != nil {
		// 	fmt.Println("[ERROR] ", err.Error())
		// 	continue
		// }

		// res := Content{
		// 	Name:    file.Name(),
		// 	Ftype:   f.ftype,
		// 	Tags:    f.GetTags().ToSlice(),
		// 	Content: b,
		// }

		// list = append(list, res)
		count -= 1
		if count == 0 {
			break
		}
	}
	wg.Wait()
	fmt.Println("[DEBUG] ", time.Since(satrtT).Seconds(), "sec")
	return list, nil
}

func asyncFilesInDir(pathFile string, list *[]Content, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	f, b, err := OpenFile(pathFile)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
	res := Content{
		Name:    f.Name,
		Ftype:   f.ftype,
		Tags:    f.GetTags().ToSlice(),
		Content: b,
	}

	m.Lock()
	*list = append(*list, res)
	m.Unlock()
}

func SearchFileWithTags(dir string, tags []string) ([]Content, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var res []Content
	pathFile, err := path.ParsePath(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		pathFile.Join(file.Name())
		f, b, err := OpenFile(pathFile.StringLinux())
		pathFile.Back()
		if err != nil {
			return nil, err
		}

		relevant := false
		for _, t := range tags {
			relevant = f.Tags.Contains(t)
			if relevant {
				break
			}
		}

		if relevant {
			res = append(res, Content{
				Name:    f.Name,
				Content: b,
				Ftype:   f.ftype,
				Tags:    f.Tags.ToSlice(),
			})
		}
	}

	return res, nil
}
