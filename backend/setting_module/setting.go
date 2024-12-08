package settingmodule

import (
	"encoding/json"
	"fmt"
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
	f, err := os.OpenFile("setting/setting.json", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return err
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(s)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
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
