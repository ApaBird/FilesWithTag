package filesmanager

import (
	"FilesWithTag/pkg/set"
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type File struct {
	Name     string
	ftype    string
	dir      string
	Tags     *set.Set
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

	if err := f.extractTags(); err != nil {
		fmt.Println("[ERROR] ", err.Error(), filePath)
		f.Tags = set.NewSet()
	}

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

func (f *File) extractTags() error {
	t := time.Now()

	content, err := os.ReadFile(f.dir + "/" + f.Name)
	if err != nil {
		return err
	}
	if time.Since(t).Seconds() > 0.1 {
		fmt.Println("[DEBUG] Long loading time", time.Since(t).Seconds(), "Size:", float64(len(content))/1024/1024, "MB")
	}
	f.Tags, f.haveTags = extractTags(content)
	return nil
}

func extractTags(content []byte) (tags *set.Set, haveTags bool) {
	tagsIndex := bytes.Index(content, []byte("Tags:"))
	tags = set.NewSet()
	if tagsIndex != -1 {
		haveTags = true
		tagsStr := string(content[tagsIndex:])
		tagsStr = tagsStr[strings.Index(tagsStr, ":")+1:]
		if strings.Contains(tagsStr, ",") {
			for _, tag := range strings.Split(tagsStr, ",") {
				if tag != "" {
					tags.Append(tag)
				}
			}
		}
	} else {
		tags = set.NewSet()
		haveTags = false
	}

	return tags, haveTags
}

func (f *File) GetTags() *set.Set {
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

	if _, err := file.WriteString(tag + ","); err != nil {
		return err
	}

	f.Tags.Append(tag)
	tags.Add(tag, f.dir+"/"+f.Name)
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

	tags.Remove(tag, f.dir+"/"+f.Name)

	return nil

}

func (f *File) HaveTage(tag string) bool {
	return f.Tags.Contains(tag)
}
