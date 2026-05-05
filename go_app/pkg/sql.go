// Just a little script for handling SQL.
// It's definitely secure

package sql

import (
	"database/sql"
	"fmt"
)

func doSQL() {
	username := "admin"
	pass := "' OR 1=1--"
    query := "SELECT * FROM users WHERE username=? AND password=?"
    db, _ := sql.Open("mysql", "user:password@/dbname")
    db.Exec(query, username, pass)
}
