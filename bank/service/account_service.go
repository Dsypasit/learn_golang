package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccoutService(accRepo repository.AccountRepository) accountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, requests NewAccountRequest) (*AccountResponse, error) {
	if requests.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}
	if strings.ToLower(requests.AccountType) != "saving" && strings.ToLower(requests.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02 15:03:05"),
		AccountType: requests.AccountType,
		Amount:      requests.Amount,
		Status:      1,
	}
	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpeted error")
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpeted error")
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			Amount:      account.Amount,
			AccountType: account.AccountType,
			Status:      account.Status,
		})
	}

	return responses, nil
}
