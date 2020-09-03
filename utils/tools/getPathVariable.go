package tools

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetPathVar(key string, r *http.Request) string {
	return mux.Vars(r)[key]
}
