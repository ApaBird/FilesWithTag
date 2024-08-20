package service

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()

	url = strings.Replace(url, "/view/", "", 1)

	file, err := os.Open("../frontend/" + url)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(body)
}
