package filesmanager

import (
	tagmap "FilesWithTag/pkg/tag_map"
	"os"
)

var tags = tagmap.NewTagMap()

func AnalyzeDir(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f := OpenFileWithTags(dir + "/" + file.Name())

		for _, tag := range f.Tags.ToSlice() {
			tags.Add(tag, dir+"/"+f.Name)
		}

	}
}
