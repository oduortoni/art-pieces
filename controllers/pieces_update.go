package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oduortoni/art-pieces/sqlite"
	types_t "github.com/oduortoni/art-pieces/types"
)

// Handler to update an existing piece
func PiecesUpdateHandler(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var piece types_t.Piece
		if err := json.NewDecoder(r.Body).Decode(&piece); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		success := sqlite.Run(sqlite.PiecesUpdate, sqlite.PiecesUpdateQuery, id, piece.Title, piece.Slug, piece.Value, piece.Description, piece.Details).(bool)
		if !success {
			http.Error(w, "Failed to update piece", http.StatusInternalServerError)
			return
		}

		// Return the updated piece
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(piece)
	}
}
