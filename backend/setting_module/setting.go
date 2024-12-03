package settingmodule

import (
	"encoding/json"
	"os"
)

type Setting struct {
	Port     string `json:"Port"`
	BasePath string `json:"BasePath"`
	LastDir  string `json:"LastDir"`
}

func NewSetting() *Setting {
	return &Setting{
		Port:     "8050",
		BasePath: "C:/",
		LastDir:  "C:/",
	}
}

func (s *Setting) Save() error {
	f, err := os.Open("setting/setting.json")
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(s)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) Change(field string, value string) error {
	switch field {
	case "Port":
		s.Port = value
	case "BasePath":
		s.BasePath = value
	case "LastDir":
		s.LastDir = value
	}

	return s.Save()
}
