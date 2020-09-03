package tools

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Parser(r *http.Request, data interface{}) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	ct := r.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(data); err != nil {
			return err
		}
	}
	return nil
}
