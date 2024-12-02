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
	Tags     *Set
	haveTags bool
}

var Ftype = map[string]string{
	"Image": ".jpg, .png, .gif, .bmp, .svg, .webp",
	"Music": ".mp3, .wav, .ogg, .aac, .m4a, .flac, .wma, .m3u",
	"Video": ".mp4, .mkv, .avi, .wmv, .flv, .mov, .webm, .mpeg",
	"Text":  ".txt, .pdf, .doc, .docx, .xls, .xlsx, .ppt, .pptx, .csv, .json",
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

// TODO при пустом "Tags:" парсит тег "", по сути не существующий тег
func extractTags(content []byte) (tags *Set, haveTags bool) {
	tagsIndex := bytes.Index(content, []byte("Tags:"))
	tags = NewSet()
	if tagsIndex != -1 {
		tagsStr := string(content[tagsIndex:])
		tags.AppendSlice(strings.Split(tagsStr[strings.Index(tagsStr, ":")+1:], ","))
		haveTags = true
	} else {
		tags = NewSet()
		haveTags = false
	}

	return tags, haveTags
}

func (f *File) GetTags() *Set {
	return f.Tags
}

func (f *File) AddTag(tag string) error {
	file, err := os.OpenFile(f.dir+"/"+f.Name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("[DEBUG]", f.Tags.ToString())

	if f.Tags.Contains(tag) {
		return nil
	}

	if !f.haveTags {
		if _, err := file.WriteString("Tags:"); err != nil {
			return err
		}
	}

	if f.Tags.Size() == 0 {
		if _, err := file.WriteString(tag); err != nil {
			return err
		}
	} else {
		if _, err := file.WriteString("," + tag); err != nil {
			return err
		}
	}
	f.Tags.Append(tag)
	return nil
}

func (f *File) RemoveTag(tag string) error {
	if !f.Tags.Contains(tag) {
		return nil
	}

	stat, err := os.Stat(f.dir + "/" + f.Name)
	if err != nil {
		return err
	}

	sizeTag := len(f.Tags.ToString())
	fmt.Println("[DEBUG]", sizeTag, f.Tags.ToString())

	if err := os.Truncate(f.dir+"/"+f.Name, stat.Size()-int64(sizeTag)); err != nil {
		return err
	}

	f.Tags.Remove(tag)

	file, err := os.OpenFile(f.dir+"/"+f.Name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("Tags:"); err != nil {
		return err
	}

	for _, t := range f.Tags.ToSlice() {
		if _, err := file.WriteString("," + t); err != nil {
			return err
		}
	}

	return nil

}
