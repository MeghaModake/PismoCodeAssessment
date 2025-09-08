package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"
)

type TransactionHandler struct {
	AccountService     *services.AccountService
	TransactionService *services.TransactionsService
}

func (t *TransactionHandler) IsValidateRequest(req datastruct.CreateTransactionRequest) error {
	if !t.AccountService.AccountExits(req.Account_ID) {
		return fmt.Errorf("Invalid Account ID %v", req.Account_ID)
	}
	// also check for amount
	//
	//req.amount
	return nil
}
func (t *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {

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
