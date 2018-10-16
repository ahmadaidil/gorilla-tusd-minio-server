package app

import (
	"net/http"

	"github.com/ahmadaidil/gorilla-tusd-minio-server/handler"
	"github.com/gorilla/mux"
)

// Handler .
func Handler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Home).Methods("GET")

	return router
}
