package database

import (
	"bizpooly/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (ac *DBContext) PostgresConnect() (*sql.DB, error) {
	connString := config.PostgresURI(ac.VConfig)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error in openning postgres connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}
