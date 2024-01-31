package database

import (
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp(sourceURL, databaseURL string) {
	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func MigrateDown(sourceURL, databaseURL string) {
	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
