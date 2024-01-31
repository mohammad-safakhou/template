package services

import (
	"context"
	"database/sql"
	"git.siz-tel.com/charging/template/internal/repository"
	"git.siz-tel.com/charging/template/internal/services/service_models"
)

type Accounts interface {
	Save(ctx context.Context, account service_models.Account) error
	Get(ctx context.Context, imsi string) (service_models.Account, error)
	GetAll(ctx context.Context) ([]service_models.Account, error)
	GetWithTXT(tx *sql.Tx) Accounts
}

type accountsService struct {
	accountRepository repository.AccountRepository
}

func (c *accountsService) Save(ctx context.Context, account service_models.Account) error {
	return c.accountRepository.InsertAccount(ctx, account)
}

func (c *accountsService) Get(ctx context.Context, imsi string) (service_models.Account, error) {
	return c.accountRepository.GetAccountByImsi(ctx, imsi)
}

func (c *accountsService) GetAll(ctx context.Context) ([]service_models.Account, error) {
	return c.accountRepository.GetAllAccounts(ctx)
}

func (c *accountsService) GetWithTXT(tx *sql.Tx) Accounts {
	c.accountRepository = c.accountRepository.GetWithTXT(tx)
	return c
}

func NewAccountsService(accountRepository repository.AccountRepository) Accounts {
	return &accountsService{
		accountRepository: accountRepository,
	}
}
