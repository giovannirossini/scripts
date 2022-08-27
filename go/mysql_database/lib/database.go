package lib

import (
	// github.com/go-sql-driver/mysql is not use in application
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Database is a pointer to DB on sqlx package
var Database *sqlx.DB

// OpenSQL is a new connection to Database
func OpenSQL() (err error) {
	err = nil
	Database, err = sqlx.Open("mysql", "golang:password@tcp(127.0.0.1:3306)/place")
	return

	err = Database.Ping()
	if err != nil {
		return
	}

	return
}
