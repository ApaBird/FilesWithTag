package filesmanager

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
)

type File struct {
	Name     string
	ftype    string
	dir      string
	Tags     []string
	haveTags bool
}

var ()

func OpenFile(filePath string) *File {
	f := File{
		Name:  path.Base(filePath),
		ftype: path.Ext(filePath),
		dir:   path.Dir(filePath),
	}

	f.extractTags()

	return &f
}

func (f *File) GetContent() ([]byte, error) {
	return f.loadFile()
}

func (f *File) loadFile() (file []byte, err error) {
	file, err = os.ReadFile(f.dir + "/" + f.Name + "." + f.ftype)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) extractTags() {
	file, err := os.Open(f.dir + "/" + f.Name + "." + f.ftype)
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	step := 100
	if stat.Size() < int64(step) {
		step = int(stat.Size())
	}

	content := []byte{}

	for i := 1; ; i++ {
		offset := stat.Size() - int64(step)*int64(i)
		if offset < 0 {
			offset = 0
		}

		sizeBuff := step
		if int64(len(content)+sizeBuff) > stat.Size() {
			sizeBuff = int(stat.Size()) - len(content)
		}

		buf := make([]byte, sizeBuff)
		_, err := file.ReadAt(buf, offset)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		content = append(content, buf[:]...)
		if bytes.Contains(content, []byte("Tags:")) {
			break
		}
	}

	tagsIndex := bytes.Index(content, []byte("Tags:"))
	if tagsIndex != -1 {
		tags := string(content[tagsIndex:])
		f.Tags = append(f.Tags, strings.Split(tags[strings.Index(tags, ":")+1:], ",")...)
		f.haveTags = true
	} else {
		f.Tags = []string{}
		f.haveTags = false
	}
}

func (f *File) GetTags() []string {
	return f.Tags
}

func (f *File) AddTag(tag string) {
	if !f.haveTags {
		os.WriteFile(f.dir+"/"+f.Name+"."+f.ftype, []byte("Tags:"), 0644)
	}

	if len(f.Tags) == 0 {
		os.WriteFile(f.dir+"/"+f.Name+"."+f.ftype, []byte(tag), 0644)
	} else {
		os.WriteFile(f.dir+"/"+f.Name+"."+f.ftype, []byte(","+tag), 0644)
	}
	f.Tags = append(f.Tags, tag)
}
