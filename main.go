package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const defaultPort = ":3000"

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Ok")
	}).Methods("GET")

	finish := make(chan bool)
	go func() {
		if err := http.ListenAndServe(":3000", router); err != nil {
			panic(err)
		}
	}()
	log.Println("Listening to http://localhost" + port)
	<-finish
}
