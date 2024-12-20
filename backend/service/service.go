package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

type ResponceError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

type Responce struct {
	Status  int    `json:"status"`
	Comment string `json:"comment"`
}

var ErrParametrs = errors.New("отсутсвуют обязательные параметры")
var ErrNotCorrectTypeParametr = errors.New("некорректный тип параметра")

func Wrapper(f func(w http.ResponseWriter, r *http.Request) any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := f(w, r)
		if resp == nil {
			return
		}
		val := reflect.ValueOf(resp)

		if val.Type().Kind() == reflect.Struct && val.IsValid() {

			w.Header().Set("Content-Type", "application/json")

			switch val.Type().Name() {
			case "ResponceError":
				w.WriteHeader((resp.(ResponceError)).Status)
			default:
				w.WriteHeader(http.StatusOK)
			}

			b, err := json.Marshal(resp)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write(b)
		} else {
			w.Write([]byte("неизвестный формат ответа сервиса"))
		}
	}
}
