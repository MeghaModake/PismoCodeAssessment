package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"
)

type TransactionHandler struct {
	AccountService     *services.AccountService
	TransactionService *services.TransactionsService
	Logger             *log.Logger
}

func (t *TransactionHandler) IsValidateRequest(req datastruct.CreateTransactionRequest) error {

	if req.Account_ID == 0 || req.OperationType_ID == 0 || req.Amount == 0 {
		return fmt.Errorf("Invalid request input, missing required %v", req)
	}
	if !t.AccountService.AccountExits(req.Account_ID) {
		return fmt.Errorf("Invalid account_id %v", req.Account_ID)
	}

	switch req.OperationType_ID {
	case 1, 2, 3: // All good
	case 4:
		if req.Amount < 0 {
			return fmt.Errorf("Ammount can not be negative for given operation_type_id %v", req.OperationType_ID)
		}
	default:
		return fmt.Errorf("Invalid operation_type_id %v", req.OperationType_ID)
	}

	return nil
}
func (t *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {

	t.Logger.Println("Received request to create transaction")
	var req datastruct.CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := t.IsValidateRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp datastruct.Transaction

	resp, err := t.TransactionService.CreateTransaction(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
