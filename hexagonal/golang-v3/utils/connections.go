package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"template/config"
)

func PostgresConnection(vConfig *viper.Viper) (*sql.DB, error) {
	connString := config.PostgresURI(vConfig)
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
