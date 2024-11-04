package db

import (
	"database/sql"

	"github.com/buonotti/apisense/v2/filesystem/locations/files"
	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

const setup = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	enabled INTEGER NOT NULL DEFAULT 1
);
`

// Setup sets up the connection to the local sqlite db. Creates the user table if it not exists
func Setup() error {
	if db != nil {
		return nil
	}

	db_, err := sql.Open("sqlite", files.DbFile())
	if err != nil {
		return err
	}

	db = db_

	_, err = db.Exec(setup)
	if err != nil {
		return err
	}

	return nil
}
