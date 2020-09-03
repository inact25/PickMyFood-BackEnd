package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	FORMART_DATE = `2006-01-02 15:04:05`
)

func DecodePathVariabel(val string, r *http.Request) string {
	param := mux.Vars(r)
	return param[val]
}

func DecodeQueryParams(val string, r *http.Request) string {
	values := r.URL.Query()
	return values[val][0]
}

func JsonDecoder(val interface{}, r *http.Request) error {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&val)
	if err != nil {
		return err
	}
	return nil
}

func GetTimeNow() string {
	return time.Now().Format(FORMART_DATE)
}
