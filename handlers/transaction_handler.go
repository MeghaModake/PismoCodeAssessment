package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	customerrors "pismo-code-assessment/customerror"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"

	"github.com/google/uuid"
)

type TransactionHandler struct {
	AccountService     *services.AccountService
	TransactionService *services.TransactionsService
	Logger             *log.Logger
}

func (t *TransactionHandler) ValidateRequest(req datastruct.CreateTransactionRequest) error {

	if req.Account_ID == 0 {
		return fmt.Errorf("account_id is required")
	}
	if req.OperationType_ID == 0 {
		return fmt.Errorf("operation_type_id is required")
	}
	if req.Amount == 0 {
		return fmt.Errorf("amount must not be zero")
	}

	if !t.AccountService.AccountExists(req.Account_ID) {
		return fmt.Errorf("%w: %d", customerrors.INVALIDACCOUNTID, req.Account_ID)
	}

	switch req.OperationType_ID {
	case datastruct.OpPurchase, datastruct.OpInstallment, datastruct.OpWithdrawal: // Amount can be postive or negative all good
	case datastruct.OpPayment:
		if req.Amount < 0 {
			return fmt.Errorf("Ammount must be positive for given operation_type_id %v", req.OperationType_ID)
		}
	default:
		return fmt.Errorf("Invalid operation_type_id %v", req.OperationType_ID)
	}

	return nil
}
func (t *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {

	requestID := uuid.New().String()

	t.Logger.Println("Received request to create transaction")
	t.Logger.Printf("[request_id=%s] handling transaction creation", requestID) // TODO :further include requestID in all logs to track

	var req datastruct.CreateTransactionRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := t.ValidateRequest(req); err != nil {
		t.Logger.Println("Validating transaction request failed!", err)
		errResp := customerrors.ErrorResponse{ErrID: http.StatusBadRequest, Errormsg: "validation failed", Details: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	var resp *datastruct.Transaction

	resp, err := t.TransactionService.Create(req)
	if err != nil {
		t.Logger.Println("Creating transaction request failed!", err)
		errResp := customerrors.ErrorResponse{ErrID: http.StatusInternalServerError, Errormsg: "failed to create transaction", Details: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
