package database

import (
	"template/transport"
)

type DBContext struct {
	*transport.ApplicationContext
}

func (ac *DBContext) RegisterDatabases() {
	psql, err := ac.PostgresConnect()
	if err != nil {
		ac.Logger.Fatal("error on connecting to postgres %v", err)
	}
	ac.PsqlDb = psql
}
