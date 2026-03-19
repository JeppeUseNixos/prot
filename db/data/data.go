package data

import (
	"database/sql"
	"fmt"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

func InitDB(dbPath string) error {
	dbConn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("could not open db: %s", err)
	}
	statement, err := dbConn.Prepare(schema)
	if err != nil {
		return fmt.Errorf("error with sql schema: %s", err)
	}

	statement.Exec()
	return nil
}

func OpenDB(dbPath string) (*sql.DB, error) {
	dbConn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %s", err)
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("db closed unexpectedly: %s", err)
	}
	return dbConn, nil
}
