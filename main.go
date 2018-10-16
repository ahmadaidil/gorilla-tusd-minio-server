package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmadaidil/gorilla-tusd-minio-server/app"
)

const defaultPort = ":3000"

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	handler := app.Handler()

	finish := make(chan bool)
	go func() {
		if err := http.ListenAndServe(":3000", handler); err != nil {
			panic(err)
		}
	}()
	log.Println("Listening to http://localhost" + port)
	<-finish
}
