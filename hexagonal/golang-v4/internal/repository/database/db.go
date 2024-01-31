package database

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func exe(db *sql.DB, tx *sql.Tx) boil.ContextExecutor {
	if tx != nil {
		return tx
	}
	return db
}
