package filesmanager

import "os"

var TagsMap = map[string][]string{}

func LoadAllTagsInDir(dir string) {
	if _, ok := TagsMap[dir]; !ok {
		TagsMap[dir] = make([]string, 0)
	}

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
			TagsMap[dir] = append(TagsMap[tag], dir+"/"+file.Name())
		}
	}
}
