package models

import (
	"main/utils"
)

func NewAccountService(accountsFile, transactionsFile string) *AccountService {
	service := &AccountService{
		AccountsFile:     accountsFile,
		TransactionsFile: transactionsFile,
	}
	service.LoadData()
	return service
}

func (s *AccountService) LoadData() {
	utils.ReadFile(s.AccountsFile, &s.Accounts)
	utils.ReadFile(s.TransactionsFile, &s.Transactions)
}

func (s *AccountService) saveData() {
	utils.WriteFile(s.AccountsFile, s.Accounts)
	utils.WriteFile(s.TransactionsFile, s.Transactions)
}

func (s *AccountService) CreateAccount(userID string) error {
	for _, acc := range s.Accounts {
		if acc.UserID == userID {
			return utils.ErrUsernameExists
		}
	}
	account := Account{UserID: userID, Balance: 0}
	s.Accounts = append(s.Accounts, account)
	s.saveData()
	return nil
}

func (s *AccountService) GetBalance(userID string) (float64, error) {
	for _, acc := range s.Accounts {
		if acc.UserID == userID {
			return acc.Balance, nil
		}
	}
	return 0, utils.ErrAccountNotFound
}

func (s *AccountService) TopUp(userID string, amount float64, description string) error {
	for i, acc := range s.Accounts {
		if acc.UserID == userID {
			s.Accounts[i].Balance += amount
			transaction := Transaction{
				ID:          len(s.Transactions) + 1,
				FromUserID:  "TOP-UP",
				ToUserID:    userID,
				Amount:      amount,
				Type:        "topup",
				Description: description,
			}
			s.Transactions = append(s.Transactions, transaction)
			s.saveData()
			return nil
		}
	}
	return utils.ErrAccountNotFound
}

func (s *AccountService) Transfer(fromUserID, toUserID string, amount float64) error {

	if fromUserID == toUserID {
		return utils.ErrSelfTransfer
	}

	var fromAcc, toAcc *Account
	for i, acc := range s.Accounts {
		if acc.UserID == fromUserID {
			fromAcc = &s.Accounts[i]
		}
		if acc.UserID == toUserID {
			toAcc = &s.Accounts[i]
		}
	}

	if fromAcc == nil || toAcc == nil {
		return utils.ErrAccountNotFound
	}

	if fromAcc.Balance < amount {
		return utils.ErrInsufficientFunds
	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount

	transaction := Transaction{
		ID:          len(s.Transactions) + 1,
		FromUserID:  fromUserID,
		ToUserID:    toUserID,
		Amount:      amount,
		Type:        "transfer",
		Description: "Transfer balance",
	}
	s.Transactions = append(s.Transactions, transaction)
	s.saveData()

	return nil
}
