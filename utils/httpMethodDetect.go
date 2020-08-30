package utils

import "net/http"

func MethodDetect(r *http.Request) string {
	switch r.Method {
	case "GET":
		return "Method Get"
	case "POST":
		return "Method POST"
	case "DELETE":
		return "Method DELETE"
	case "PUT":
		return "Method PUT"
	default:
		return "Method not yet implemented"
	}

}
