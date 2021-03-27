package database

import (
	"backend-service/transport"
)

type DBContext struct {
	*transport.ApplicationContext
}

func (ac *DBContext) RegisterDatabases() {
	psql, err := ac.PostgresConnect()
	if err != nil {
		ac.Logger.Fatal("error on connecting to postgres %v", err)
	}
	redis, err := ac.RedisConnect()
	if err != nil {
		ac.Logger.Fatal("error on connecting to redis %v", err)
	}
	ac.PsqlDb = psql
	ac.Redis = redis
}
