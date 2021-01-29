package main

import (
	"fmt"
	"database/sql"
    _ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

var (
	dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
	dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
	instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
)

// InitializeDatabaseConnection initializes database connection
func InitializeDatabaseConnection() (*sql.DB, error) {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
			socketDir = "/cloudsql"
	}
	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}