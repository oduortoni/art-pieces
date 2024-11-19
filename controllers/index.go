package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found!", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s", "Hello")
}
