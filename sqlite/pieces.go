package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	log_t "github.com/oduortoni/art-pieces/log"
	types_t "github.com/oduortoni/art-pieces/types"
)

/*
* To be called only once, in order t initialize the pieces table in the database
 */
func CreatePiecesTable(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		_, err := db.Exec(query)
		return err == nil
	}
	return false
}

/*
* create a piece
* args: title, slug, value, description, details
 */
 func PieceCreate(db *sql.DB, args ...any) any {
	errorStr := ""

	query, ok := args[0].(string)
	if !ok {
		errorStr += "query"
	}
	title, ok := args[1].(string)
	if !ok {
		errorStr += "title"
	}
	slug, ok := args[2].(string)
	if !ok {
		errorStr += "slug"
	}
	value, ok := args[3].(float64)
	if !ok {
		errorStr += "value"
	}
	description, ok := args[4].(string)
	if !ok {
		errorStr += "description"
	}
	details, ok := args[5].(string)
	if !ok {
		errorStr += "details"
	}

	if len(errorStr) != 0 {
		log_t.LogW("pieces-create", errorStr, nil)
		return false
	}
	_, err := db.Exec(query, title, slug, value, description, details)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err == nil
}


/*
* select all pieces from the database
 */
func PiecesSelect(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pieces := []*types_t.Piece{}
	for rows.Next() {
		var id int
		var title string
		var description string
		if err := rows.Scan(&id, &title, &description); err != nil {
			log.Fatal(err)
		}
		piece := &types_t.Piece{
			Identifier:  id,
			Title:       title,
			Description: description,
		}
		pieces = append(pieces, piece)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return pieces
}

/*
* select a piece from the database by id
 */
func PiecesSelectById(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if !ok {
		return nil
	}
	pieceId, ok := args[1].(int)
	if !ok {
		return nil
	}
	rows, err := db.Query(query, pieceId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pieces := []*types_t.Piece{}
	for rows.Next() {
		var id int
		var title string
		var slug string
		var value float64
		var description string
		var details string
		if err := rows.Scan(&id, &title, &slug, &value, &description, &details); err != nil {
			log.Fatal(err)
		}
		piece := &types_t.Piece{
			Identifier:  id,
			Title:       title,
			Slug:        slug,
			Value:       value,
			Description: description,
			Details:     details,
		}
		pieces = append(pieces, piece)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return pieces
}

/* updates a single piece on the database */
func PiecesUpdate(db *sql.DB, args ...any) any {
	errorStr := ""

	query, ok := args[0].(string)
	if !ok {
		errorStr += "query"
	}
	id, ok := args[1].(int)
	if !ok {
		errorStr += "id"
	}
	title, ok := args[2].(string)
	if !ok {
		errorStr += "title"
	}
	slug, ok := args[3].(string)
	if !ok {
		errorStr += "slug"
	}
	value, ok := args[4].(float64)
	if !ok {
		errorStr += "value"
	}
	description, ok := args[5].(string)
	if !ok {
		errorStr += "description"
	}
	details, ok := args[6].(string)
	if !ok {
		errorStr += "details"
	}

	if len(errorStr) != 0 {
		log_t.LogW("pieces-create", errorStr, nil)
		return false
	}
	_, err := db.Exec(query, title, slug, value, description, details, id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return err == nil
}

/*
* deletes an entire piece
 */
func PiecesDelete(db *sql.DB, args ...any) any {
	query, ok := args[0].(string)
	if ok {
		pieceId, ok := args[1].(int)
		if ok {
			_, err := db.Exec(query, pieceId)
			return err == nil
		}
	}
	return false
}
