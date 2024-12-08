package filesmanager

import (
	"FilesWithTag/pkg/set"
	"fmt"
	"sync"
	"time"
)

func asyncFilesInDir(pathFile string, list *[]Content, tagsInDir *set.Set, wg *sync.WaitGroup, m *sync.Mutex) {
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
	tagsInDir.AppendSlice(f.GetTags().ToSlice())
	*list = append(*list, res)
	m.Unlock()
}

func asyncLoadFilesList(files []string) []Content {
	var list []Content
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	start := time.Now()
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			f, b, err := OpenFile(file)
			if err != nil {
				fmt.Println("[ERROR] ", err.Error())
			}

			m.Lock()
			list = append(list, Content{
				Name:    f.Name,
				Ftype:   f.ftype,
				Tags:    f.GetTags().ToSlice(),
				Content: b,
			})
			m.Unlock()
		}(file)
	}
	wg.Wait()
	if time.Since(start).Seconds() > 0.1 {
		fmt.Println("Time asyncLoadFilesList:", time.Since(start).Seconds())
	}
	return list
}
