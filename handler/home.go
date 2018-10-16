package handler

import (
	"fmt"
	"net/http"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok")
}
