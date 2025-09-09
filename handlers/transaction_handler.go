package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	customerrors "pismo-code-assessment/CustomErrors"
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
		return fmt.Errorf(customerrors.INVALIDACCOUNTID, req.Account_ID)
	}

	switch req.OperationType_ID {
	case 1, 2, 3: // Amount can be postive or negative all good
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
