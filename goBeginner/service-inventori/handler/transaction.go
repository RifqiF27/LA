package handler

import (
	"database/sql"
	"service-inventory/model"
	"service-inventory/repository"
	"service-inventory/service"
	"service-inventory/utils"
)

func AddTransaction(db *sql.DB) {
	_, valid := utils.SessionAdmin()
	if !valid {
		return
	}
	transaction := model.Transaction{}
	err := utils.DecodeJSONFile("body.json", &transaction)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	transactionService := service.NewTransactionService(repository.NewTransactionRepo(db))
	newTransaction, err := transactionService.CreateTransactionService(&transaction)
	if err != nil {
		utils.SendJSONResponse(400, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(201, "transaction created successfully", newTransaction)
}

func GetTransactions(db *sql.DB) {
	_, valid := utils.Session()
	if !valid {
		return
	}
    var history model.TransactionHistory
    err := utils.DecodeJSONFile("body.json", &history)
    if err != nil {
        utils.SendJSONResponse(400, err.Error(), nil)
        return
    }

    transactionService := service.NewTransactionService(repository.NewTransactionRepo(db))
    transactions, err := transactionService.GetTransactionsService(history.ItemName)
    if err != nil {
        utils.SendJSONResponse(500, err.Error(), nil)
        return
    }

    utils.SendJSONResponse(200, "transactions retrieved successfully", transactions)
}
