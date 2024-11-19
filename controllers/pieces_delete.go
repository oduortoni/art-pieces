package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oduortoni/art-pieces/sqlite"
	types_t "github.com/oduortoni/art-pieces/types"
)

// Handler to delete a piece
func PiecesDeleteHandler(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		success := sqlite.Run(sqlite.PiecesDelete, sqlite.PiecesDeleteQuery, id).(bool)
		if !success {
			http.Error(w, "Failed to delete piece", http.StatusInternalServerError)
			return
		}

		// Return a 204 No Content status
		w.WriteHeader(http.StatusNoContent)
	}
}
