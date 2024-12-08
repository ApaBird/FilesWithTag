package filesmanager

import (
	"FilesWithTag/pkg/path"
	tagmap "FilesWithTag/pkg/tag_map"
	"fmt"
	"os"
	"sync"
	"time"
)

var tags = tagmap.NewTagMap()

func AnalyzeStorage(dir string) {
	wg := &sync.WaitGroup{}
	start := time.Now()
	loadFile := 0
	p, err := path.ParsePath(dir)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
	analyzeDir(p, &sync.Mutex{}, wg, &loadFile)
	wg.Wait()
	fmt.Println("Time analyze:", time.Since(start).Seconds())
	fmt.Println("Files Loaded: ", loadFile)
}

func analyzeDir(dir *path.Path, m *sync.Mutex, wg *sync.WaitGroup, loadfile *int) {
	wg.Add(1)
	defer wg.Done()
	files, err := os.ReadDir(dir.StringLinux())
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			p, err := path.ParsePath(dir.StringLinux() + "/" + file.Name())
			if err != nil {
				fmt.Println("[ERROR] ", err.Error())
				return
			}
			go analyzeDir(p, m, wg, loadfile)
			continue
		}

		m.Lock()
		*loadfile += 1
		m.Unlock()

		dir.Join(file.Name())
		f := OpenFileWithTags(dir.StringLinux())

		for _, tag := range f.Tags.ToSlice() {
			m.Lock()
			tags.Add(tag, dir.StringLinux())
			m.Unlock()
		}
		dir.Back()

	}
}
