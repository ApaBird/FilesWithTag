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

func OpenFile(filePath string) (file *File, content []byte, err error) {
	filePath = strings.Replace(filePath, "\\", "/", -1)

	f := File{
		Name:  path.Base(filePath),
		ftype: path.Ext(filePath),
		dir:   path.Dir(filePath),
	}

	b, err := f.loadFile()
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return nil, nil, err
	}

	f.Tags, f.haveTags = extractTags(b)

	return &f, b, nil

}

func OpenFileWithTags(filePath string) *File {
	filePath = strings.Replace(filePath, "\\", "/", -1)
	f := File{
		Name:  path.Base(filePath),
		ftype: path.Ext(filePath),
		dir:   path.Dir(filePath),
	}

	fmt.Println("[DEBUG]", f.dir)
	fmt.Println("[DEBUG]", f.Name)

	f.extractTags()

	return &f
}

func (f *File) GetContent() ([]byte, error) {
	return f.loadFile()
}

func (f *File) loadFile() (file []byte, err error) {
	file, err = os.ReadFile(f.dir + "/" + f.Name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) extractTags() {
	fmt.Println("[DEBUG]", f.dir+"/"+f.Name)
	file, err := os.Open(f.dir + "/" + f.Name)
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

	content := make([]byte, 0)

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
		fmt.Println(stat.Size(), sizeBuff, offset, len(content), step)

		content = append(content, buf[:]...)
		if bytes.Contains(content, []byte("Tags:")) {
			break
		}

		if len(content) >= int(stat.Size()) {
			break
		}
		//time.Sleep(time.Second * 30)
	}

	f.Tags, f.haveTags = extractTags(content)
}

func extractTags(content []byte) (tags []string, haveTags bool) {
	tagsIndex := bytes.Index(content, []byte("Tags:"))
	if tagsIndex != -1 {
		tagsStr := string(content[tagsIndex:])
		tags = append(tags, strings.Split(tagsStr[strings.Index(tagsStr, ":")+1:], ",")...)
		haveTags = true
	} else {
		tags = []string{}
		haveTags = false
	}

	return tags, haveTags
}

func (f *File) GetTags() []string {
	return f.Tags
}

func (f *File) AddTag(tag string) error {
	file, err := os.OpenFile(f.dir+"/"+f.Name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if !f.haveTags {
		if _, err := file.WriteString("Tags:"); err != nil {
			return err
		}
	}

	if len(f.Tags) == 0 {
		if _, err := file.WriteString(tag); err != nil {
			return err
		}
	} else {
		if _, err := file.WriteString("," + tag); err != nil {
			return err
		}
	}
	f.Tags = append(f.Tags, tag)
	return nil
}
