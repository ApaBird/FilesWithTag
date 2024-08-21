package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	fmt.Println(url)

	url = strings.Replace(url, "/view/", "", 1)

	_, err := os.Stat("../frontend/" + url)

	check := !os.IsNotExist(err)

	if check {
		http.ServeFile(w, r, "../frontend/"+url)
	} else {
		w.Write([]byte("Пошел нахуй"))
	}

}
