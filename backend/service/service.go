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

var ErrParametrs = errors.New("Отсутсвуют обязательные параметры")

func Wrapper(f func(w http.ResponseWriter, r *http.Request) any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := f(w, r)
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
