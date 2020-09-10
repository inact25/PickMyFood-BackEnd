package configs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	port := "8080"
	fmt.Println("Starting Web Server at port: " + port)
	err := http.ListenAndServe(": "+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
