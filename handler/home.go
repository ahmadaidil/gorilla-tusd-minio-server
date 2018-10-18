package handler

import (
	"net/http"

	"github.com/ahmadaidil/gorilla-tusd-minio-server/view"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	view.Home(w, r)
}
