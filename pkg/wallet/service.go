package wallet

import (
	"errors"

	"github.com/surush3005/Wallet/pkg/types"
)

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, errors.New("Phone already registered")
		}
	}

	s.nextAccountID++

	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}

	s.accounts = append(s.accounts, account)

	return account, nil
}

func (s *Service) Deposit(AccountId int64, amount types.Money) {
	if amount <= 0 {

		return
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == AccountId {
			account = acc
			break
		}
	}

	if account == nil {
		return
	}

	account.Balance += amount
}

func (s *Service) Reject(paymentID string) error {
	var targetPayment *types.Payment
	for _, payment := range s.payments{
		if payment.ID = paymentID{
			targetPayment = payment
		}
	}
	
	if targetPayment == nil {
		return ErrPaymentNotFound
	}

	var targetAccount *types.Account
	for _, account := range s.accounts{
		if account.ID == targetPayment.AcoountID{
			targetAccount = account
			break
		}
	}

	if targetAccount = nil {
		return ErrAccountNotFound

	}

	targetPayment.Status = types.PaymentStatusFail
	targetAccount.Balance += targetPayment.Amount
}
