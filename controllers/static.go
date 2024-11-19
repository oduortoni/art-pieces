package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

func Static(w http.ResponseWriter, r *http.Request) {
	spath := r.URL.Path

	path := "." + spath

	stat, err := os.Stat(path)
	if err != nil {
		// file does not exist
		http.Error(w, "File does not exist", http.StatusNotFound)
		return
	}

	if stat.IsDir() {
		// forbid access to directories
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// determine MIME type based on the file extension
	ext := filepath.Ext(path)
	mime := "application/octet-stream"
	switch ext {
	case ".css":
		mime = "text/css"
	case ".js":
		mime = "application/javascript"
	default:
		mime = "application/octet-stream"
	}

	w.Header().Set("Content-Type", mime)
	http.ServeFile(w, r, path)
}
