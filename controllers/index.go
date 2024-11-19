package controllers

import (
	"net/http"

	types_t "github.com/oduortoni/art-pieces/types"
)

func Index(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Page not found!", http.StatusNotFound)
			return
		}

		http.ServeFile(w, r, "templates/index.html")
	}
}
