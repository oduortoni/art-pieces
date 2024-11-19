package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oduortoni/art-pieces/sqlite"
	types_t "github.com/oduortoni/art-pieces/types"
)

// Handler to get all pieces
func PiecesSelectHandler(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pieces := sqlite.Run(sqlite.PiecesSelect, sqlite.PiecesSelectAllQuery)
		if pieces == nil {
			http.Error(w, "No pieces found", http.StatusNotFound)
			return
		}

		// Return the list of pieces as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pieces)
	}
}

// Handler to get a piece by ID
func PiecesSelectByIdHandler(container *types_t.Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		pieces := sqlite.Run(sqlite.PiecesSelectById, sqlite.PiecesSelectByIdQuery, id).([]*types_t.Piece)
		if len(pieces) == 0 {
			http.Error(w, "Piece not found", http.StatusNotFound)
			return
		}

		// Return the piece as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pieces[0])
	}
}
