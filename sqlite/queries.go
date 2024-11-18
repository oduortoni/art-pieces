package sqlite

const  (
	PiecesCreateTableQuery = "CREATE TABLE IF NOT EXISTS pieces (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, slug TEXT, value REAL, description TEXT, details TEXT);"
	/* CREATE*/ PiecesInsertQuery = `INSERT INTO pieces(title, slug, value, description, details) VALUES (?, ?, ?, ?, ?)`
	/* READ ALL */ PiecesSelectAllQuery = "SELECT id, title, description FROM pieces"
	/* READ ONE */ PiecesSelectByIdQuery = "SELECT id, title, slug, value, description, details FROM pieces WHERE id = ?"
	/* UPDATE */ PiecesUpdateQuery = `UPDATE pieces SET title = ?, slug = ?, value = ?, description = ?, details = ? WHERE id = ?`
	/* DELETE */ PiecesDeleteQuery = `DELETE FROM pieces WHERE id = ?`
)
