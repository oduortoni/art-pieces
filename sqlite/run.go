package sqlite

import (
	"database/sql"
	"log"
)

/*
* used to open a database connection and invoking the callback that does the real work
 */
func Run(callback func(db *sql.DB, args ...any) any, args ...any) any {
	db, err := sql.Open("sqlite3", "resources/database/art.db") // open or create if it doesnt exist
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return callback(db, args...)
}
