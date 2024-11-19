package main

import (
	"fmt"
	"net/http"

	"github.com/oduortoni/art-pieces/controllers"
	utils "github.com/oduortoni/art-pieces/lib"
	sqlite_t "github.com/oduortoni/art-pieces/sqlite"
	types_t "github.com/oduortoni/art-pieces/types"

	"github.com/gorilla/mux"
)

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	sqlite_t.Run(sqlite_t.CreatePiecesTable, sqlite_t.PiecesCreateTableQuery) // initialize the table in the database

	var container types_t.Container
	router := mux.NewRouter()
	router.HandleFunc("/pieces", controllers.PiecesSelectHandler(&container)).Methods("GET")
	router.HandleFunc("/pieces", controllers.PieceCreateHandler(&container)).Methods("POST")
	router.HandleFunc("/pieces/{id:[0-9]+}", controllers.PiecesSelectByIdHandler(&container)).Methods("GET")
	router.HandleFunc("/pieces/{id:[0-9]+}", controllers.PiecesUpdateHandler(&container)).Methods("PUT")
	router.HandleFunc("/pieces/{id:[0-9]+}", controllers.PiecesDeleteHandler(&container)).Methods("DELETE")

	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(portStr, router)
}
