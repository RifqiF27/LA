package service

import (
	"travelika/model"
	"travelika/repository"
)

type TransactionService struct {
	Repo repository.TransactionsRepository
}

func NewTransactionService(repo repository.TransactionsRepository) *TransactionService {
	return &TransactionService{Repo: repo}
}

func (s *TransactionService) CreateTransaction(transaction model.Transaction) error {
	return s.Repo.Create(transaction)
}