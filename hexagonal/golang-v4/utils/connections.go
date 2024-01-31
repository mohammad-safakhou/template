package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func PostgresConnection(host, port, user, pass, database, sslmode string, maxOpenConns, maxIdleConns int, timeout time.Duration) (*sql.DB, error) {
	connString := PostgresURI(host, port, user, pass, database, sslmode)
	log.Println("postgres options -> " + connString)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error in openning postgres connection: %w", err)
	}

	conn.SetMaxOpenConns(maxOpenConns)
	conn.SetMaxIdleConns(maxIdleConns)

	dbContext, _ := context.WithTimeout(context.Background(), timeout)
	err = conn.PingContext(dbContext)
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func PostgresURI(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, database, sslmode)
}

func CreateRedisConnection(ctx context.Context, host, port string, db int, timeout time.Duration) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", host, port),
		DialTimeout: timeout,
		DB:          db,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	if pong != "PONG" {
		return nil, fmt.Errorf("expected PONG, got %s", pong)
	}

	return client, nil
}

func PostgresUrl(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		pass,
		host,
		port,
		database,
		sslmode,
	)
}
