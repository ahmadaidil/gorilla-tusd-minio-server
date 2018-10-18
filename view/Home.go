package view

import (
	"log"
	"net/http"
)

// Home view
func Home(w http.ResponseWriter, r *http.Request) {
	if err := render.HTML(w, http.StatusOK, "home", nil); err != nil {
		log.Fatal(err)
	}
}
