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
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, pass)
	db, _ := sql.Open("mysql", "user:password@/dbname")
	db.Exec(query)
}
