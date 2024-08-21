package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configstruct struct {
	Port string `json:"port"`
}

func Readconfig(path string) (config Configstruct) {
	data, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	bytedata, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(bytedata, &config)
	if err != nil {
		fmt.Println(err.Error())
	}

	return config
}
