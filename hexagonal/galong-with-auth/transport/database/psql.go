package database

import (
	"backend-service/config"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
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

func (ac *DBContext) RedisConnect() (*redis.Client ,error) {
	redisHost := ac.VConfig.GetString("database.redis.host")
	redisPort := ac.VConfig.GetString("database.redis.port")
	redisDb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       1,
	})
	_, err := redisDb.Ping(context.TODO()).Result()
	if err != nil {
		return nil, fmt.Errorf("error in pinging redis database: %w", err)
	}
	redisDb.FlushAll(context.TODO())
	return redisDb, nil
}