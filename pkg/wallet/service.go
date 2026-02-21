package wallet

import "github.com/surush3005/Wallet/pkg/types"

type Service struct {
	accounts []types.Amountcont
	payments []types.Payment
}

func RegisterAccount(service *Service, phone types.Phone) {
	for _, accounts := range service.accounts {
		if accounts.Phone == phone {
			return
		}
	}

	service.nextAccountId++

	service.accounts = append(service.accounts, types.Account{
		ID:      service.nextAccountID,
		Phone:   phone,
		Balance: 0,
	})
}
