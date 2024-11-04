package service

import (
	"errors"
	"service-inventory/model"
	"service-inventory/repository"
)

type TransactionService struct {
    Repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
    return &TransactionService{Repo: repo}
}

func (ts *TransactionService) CreateTransactionService(transaction *model.Transaction) (*model.Transaction, error) {
    if transaction.TransactionType != "in" && transaction.TransactionType != "out" {
        return nil, errors.New("transaction type must be 'in' or 'out'")
    }
    if transaction.ItemID == 0 {
		return nil, errors.New("item ID is required")
	}
	if transaction.UserID == nil {
		return nil, errors.New("user ID is required")
	}
    if transaction.Quantity <= 0 {
        return nil, errors.New("quantity must be greater than zero")
    }
    return ts.Repo.CreateTransaction(transaction)
}

func (ts *TransactionService) GetTransactionsService(itemName string) ([]model.TransactionHistory, error) {
    return ts.Repo.GetTransactions(itemName)
}
