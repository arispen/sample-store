package main

import (
	"database/sql"
    _ "github.com/jackc/pgx/v4/stdlib"
)

var (
	dbUrl                = mustGetenv("DB_URL")                 // e.g. 'my-database'
)

// InitializeDatabaseConnection initializes database connection
func InitializeDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}