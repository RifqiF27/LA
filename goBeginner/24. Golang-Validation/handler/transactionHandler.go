package handler

import (
	"encoding/json"
	"net/http"
	"travelika/helper"
	"travelika/model"
	"travelika/service"
	"travelika/utils"

	"go.uber.org/zap"
)

type TransactionHandler struct {
	Service   *service.TransactionService
	Log       *zap.Logger
	validator *helper.Validator
}

func NewTransactionHandler(service *service.TransactionService, logger *zap.Logger) *TransactionHandler {
	return &TransactionHandler{
		Service:   service,
		Log:       logger,
		validator: helper.NewValidator(),
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		h.Log.Error("Handler: invalid request payload", zap.Error(err))
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "invalid request payload", nil)
		return
	}

	if err := h.validator.ValidateStruct(transaction); err != nil {
		formattedError := helper.FormatValidationError(err)
		h.Log.Error("Handler: validation failed", zap.String("error", formattedError))
		utils.SendJSONResponse(w, false, http.StatusBadRequest, formattedError, nil)
		return
	}

	switch transaction.Status {
	case "ok":
		transaction.StatusTrx = true
	case "cancel":
		transaction.StatusTrx = false
	default:
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "invalid status value", nil)
		return
	}

	err := h.Service.CreateTransaction(transaction)
	if err != nil {
		h.Log.Error("Handler: invalid created transaction", zap.Error(err))
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, "invalid created transaction", nil)
		return
	}
	message := "transaction success"
	if !transaction.StatusTrx {
		message = "transaction canceled"
	}
	utils.SendJSONResponse(w, true, http.StatusOK, message, nil)

}
