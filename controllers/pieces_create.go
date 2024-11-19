package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oduortoni/art-pieces/sqlite"
	types_t "github.com/oduortoni/art-pieces/types"
)

// Handler to create a new piece
func PieceCreateHandler(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var piece types_t.Piece
		// Decode the incoming JSON to a Piece struct
		if err := json.NewDecoder(r.Body).Decode(&piece); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Query to insert a new piece
		success := sqlite.Run(sqlite.PieceCreate, sqlite.PiecesInsertQuery, piece.Title, piece.Slug, piece.Value, piece.Description, piece.Details).(bool)
		if !success {
			http.Error(w, "Failed to create piece", http.StatusInternalServerError)
			return
		}

		// Return a 201 Created status
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(piece)
	}
}
