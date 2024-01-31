package repository

import (
	"context"
	"database/sql"
	"git.siz-tel.com/charging/template/internal/services/service_models"
)

type AccountRepository interface {
	GetAllAccounts(ctx context.Context) ([]service_models.Account, error)
	GetAccountByImsi(ctx context.Context, imsi string) (service_models.Account, error)
	InsertAccount(ctx context.Context, sma service_models.Account) error
	GetWithTXT(tx *sql.Tx) AccountRepository
}
