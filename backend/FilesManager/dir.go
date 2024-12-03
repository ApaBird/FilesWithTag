package filesmanager

import (
	"fmt"
	"os"
	"strings"
)

type Dir struct {
	Name    string  `json:"name"`
	Dir     string  `json:"dir"`
	Content *[]*Dir `json:"content"`
}

func NewDir(name string, dir string) Dir {
	content := make([]*Dir, 0)
	return Dir{
		Name:    name,
		Dir:     dir,
		Content: &content,
	}
}

func (d *Dir) AddDir(dir Dir) {
	content := *d.Content
	content = append(content, &dir)
	d.Content = &content
}

func (d *Dir) AddDirByName(dir string) {
	content := *d.Content
	con := make([]*Dir, 0)
	content = append(content, &Dir{
		Name:    dir,
		Dir:     d.Dir + "/" + dir,
		Content: &con,
	})

	d.Content = &content
}

// func (d *Dir) AddDirByName(dir string) {
// 	dir = strings.Trim(dir, "/")
// 	split := strings.Split(dir, "/")

// 	if split[0] == d.Name {
// 		if len(split) == 1 {
// 			return
// 		}
// 		split = split[1:]
// 	} else {
// 		return
// 	}

// 	for _, name := range split{

// 	}
// }

func (d *Dir) FindDir(path string) *Dir {
	// fmt.Println(path)
	// fmt.Println(d.Name)
	path = strings.Trim(path, "/")
	path = strings.ReplaceAll(path, "//", "/")
	split := strings.Split(path, "/")
	fmt.Println("[DEBUG]", "Find Dir ", split)
	if split[0] == d.Name {
		// fmt.Println("YES")
		if len(split) == 1 {
			return d
		}
		split = split[1:]
	} else {
		return nil
	}

	for _, dir := range *d.Content {
		if dir.Name == split[0] {
			// fmt.Println("yes")
			// fmt.Println(*dir.Content)
			if len(split) == 1 {
				// fmt.Println("finded")
				return dir
			}
			return dir.FindDir(strings.Join(split, "/"))
		}
	}
	return nil
}

func GetDirs(path string) []string {
	var dirs []string

	path = strings.Trim(path, "/")
	path = strings.ReplaceAll(path, "//", "/")

	files, err := os.ReadDir(path)
	if err != nil {
		return dirs
	}
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, strings.Replace(path+"/"+f.Name(), "\\", "/", -1))
		}
	}
	return dirs
}
