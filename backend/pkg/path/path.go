package path

import (
	"errors"
	"runtime"
	"strings"
)

type Path struct {
	Dirs []string
}

var (
	ErrDubleSlash = errors.New("double slash")
)

func NewPath(dirs []string) *Path {
	return &Path{Dirs: dirs}
}

func ParsePath(path string) (*Path, error) {
	slash := strings.Contains(path, "/")
	backSlash := strings.Contains(path, "\\")
	if slash && backSlash {
		return nil, ErrDubleSlash
	}

	sep := ""
	if slash {
		sep = "/"
	} else {
		sep = "\\"
	}

	path = strings.Trim(path, sep)
	path = strings.ReplaceAll(path, sep+sep, sep)

	return NewPath(strings.Split(path, sep)), nil
}

func (p *Path) String() string {
	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}

	if len(p.Dirs) == 1 {
		return p.Dirs[0] + sep
	}
	return strings.Join(p.Dirs, sep)
}

func (p *Path) StringLinux() string {
	return strings.Join(p.Dirs, "/")
}

func (p *Path) StringWindows() string {
	return strings.Join(p.Dirs, "\\")
}

func (p *Path) Join(dir string) {
	p.Dirs = append(p.Dirs, dir)
}

func (p *Path) Back() {
	p.Dirs = p.Dirs[:len(p.Dirs)-1]
}

func (p *Path) Ext() string {
	fileName := p.Dirs[len(p.Dirs)-1]
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return ""
	}
	return fileName[dotIndex:]
}
