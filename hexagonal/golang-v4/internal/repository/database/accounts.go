package database

import (
	"context"
	"database/sql"
	"errors"

	"git.siz-tel.com/charging/template/internal/repository"
	boiler_models "git.siz-tel.com/charging/template/internal/repository/boiler"
	"git.siz-tel.com/charging/template/internal/services/service_models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type accountRepo struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
	tx      *sql.Tx
}

func NewAccountRepo(dbRead, dbWrite *sql.DB) repository.AccountRepository {
	return &accountRepo{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

func (a *accountRepo) GetAllAccounts(ctx context.Context) (accounts []service_models.Account, err error) {
	allAccounts, err := boiler_models.Accounts().All(ctx, exe(a.dbRead, a.tx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service_models.ErrRecordNotFound
		}
		return
	}

	return convertToServiceAccounts(allAccounts), nil
}

func (a *accountRepo) GetAccountByImsi(ctx context.Context, imsi string) (account service_models.Account, err error) {
	anAccount, err := boiler_models.Accounts(qm.Where("imsi=?", imsi)).One(ctx, exe(a.dbRead, a.tx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service_models.Account{}, service_models.ErrRecordNotFound
		}
		return
	}

	sliceOfAnAccount := boiler_models.AccountSlice{anAccount}
	return convertToServiceAccounts(sliceOfAnAccount)[0], nil
}

func (a *accountRepo) InsertAccount(ctx context.Context, sma service_models.Account) error {
	var account boiler_models.Account
	account.Imsi = sma.IMSI
	if err := account.Insert(ctx, exe(a.dbWrite, a.tx), boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (a *accountRepo) GetWithTXT(tx *sql.Tx) repository.AccountRepository {
	a.tx = tx
	return a
}
