// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	fmt.Printf("Server listening on %s\n", port)
// 	http.ListenAndServe(":9000", nil)
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {

//		}
//	}
package main

import (
	// "fmt"

	"fmt"

	sqlite_t "github.com/oduortoni/art-pieces/sqlite"
	type_t "github.com/oduortoni/art-pieces/types"
)

func main() {
	sqlite_t.Run(sqlite_t.CreatePiecesTable, sqlite_t.PiecesCreateTableQuery) // initialize the table in the database
	sqlite_t.Run(sqlite_t.PieceCreate, sqlite_t.PiecesInsertQuery, "The Piece", "the-piece", 0.0, "The piece is the most valuable item ever since creation. It has been, is and always will be available. It has no intrinsic value attached to it. It is literary priceless. The piece was first discovered at the inception of the universe. It is described as having no form but existing in all forms all the time", "Year: 0-.-Locaton: void-.-Comprises: all-.-Culture: all-.-Size: Infinite")
	piecesSelect, ok := sqlite_t.Run(sqlite_t.PiecesSelect, sqlite_t.PiecesSelectAllQuery).([]*type_t.Piece)
	if ok {
		for _, piece := range piecesSelect {
			fmt.Println(piece)
		}
	}

	piecesUpdate, ok := sqlite_t.Run(sqlite_t.PiecesSelect, sqlite_t.PiecesSelectAllQuery, 2).([]*type_t.Piece)
	if ok {
		fmt.Println(piecesUpdate[0])
	}

	sqlite_t.Run(sqlite_t.PiecesUpdate, sqlite_t.PiecesUpdateQuery, 2, "Piece numero dos", "the-dos", 90.0, "Nummero dos esta dos", "Year: 1903-.-Locaton: Argentina-.-Comprises: carbon-.-Culture: latina-.-Size: 100x50in")

	/* list all pieces before deleting */
	piecesDel, ok := sqlite_t.Run(sqlite_t.PiecesSelect, sqlite_t.PiecesSelectAllQuery).([]*type_t.Piece)
	if ok {
		for _, piece := range piecesDel {
			fmt.Println(piece)
		}
	}
	fmt.Printf("\n\n\n")
	sqlite_t.Run(sqlite_t.PiecesDelete, sqlite_t.PiecesDeleteQuery, 2)
	sqlite_t.Run(sqlite_t.PiecesDelete, sqlite_t.PiecesDeleteQuery, 3)
	sqlite_t.Run(sqlite_t.PiecesDelete, sqlite_t.PiecesDeleteQuery, 4)
	sqlite_t.Run(sqlite_t.PiecesDelete, sqlite_t.PiecesDeleteQuery, 5)
	piecesDel, ok = sqlite_t.Run(sqlite_t.PiecesSelect, sqlite_t.PiecesSelectAllQuery).([]*type_t.Piece)
	if ok {
		for _, piece := range piecesDel {
			fmt.Println(piece)
		}
	}
}
