package database

import (
	boiler_models "git.siz-tel.com/charging/template/internal/repository/boiler"
	"git.siz-tel.com/charging/template/internal/services/service_models"
)

func convertToServiceAccounts(bma boiler_models.AccountSlice) []service_models.Account {

	serviceAccounts := make([]service_models.Account, len(bma))

	for i, account := range bma {
		serviceAccounts[i] = service_models.Account{
			IMSI:      account.Imsi,
			CreatedAt: account.CreatedAt,
		}
	}

	return serviceAccounts
}
