package settingmodule

import (
	"encoding/json"
	"os"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var appSetting *Setting

func Init() error {
	ex, err := exists("setting")
	if err != nil {
		return err
	}
	if !ex {
		if err := os.Mkdir("setting", 0777); err != nil {
			return err
		}
	}

	ex, err = exists("setting/setting.json")
	if err != nil {
		return err
	}

	if !ex {
		f, err := os.Create("setting/setting.json")
		if err != nil {
			return err
		}
		appSetting = NewSetting()
		enc := json.NewEncoder(f)

		err = enc.Encode(appSetting)
		if err != nil {
			return err
		}

		f.Close()
	} else {
		f, err := os.Open("setting/setting.json")
		if err != nil {
			return err
		}

		dec := json.NewDecoder(f)
		err = dec.Decode(&appSetting)
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}

func GetSetting() *Setting {
	return appSetting
}
