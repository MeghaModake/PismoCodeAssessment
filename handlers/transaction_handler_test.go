package handlers

import (
	"log"
	"os"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"
	"testing"
)

var ts *TransactionHandler

func IsValidateRequestSub(inputAccountID int, inputOperationTypeID int, inputAmount float64) bool {

	var req datastruct.CreateTransactionRequest
	req.Account_ID = inputAccountID
	req.OperationType_ID = inputOperationTypeID
	req.Amount = inputAmount
	if err := ts.IsValidateRequest(req); err != nil {
		return false
	}

	return true

}
func setup() {
	logger := log.New(os.Stdout, "APP_LOG: ", log.LstdFlags|log.Lshortfile)
	accountService := services.NewAccountService(logger)
	transactionService := services.NewTransactionsService(logger)
	ts = &TransactionHandler{AccountService: accountService, TransactionService: transactionService, Logger: logger}

	// create account with Document_Number = 101 which will have account id = 1
	var req datastruct.CreateAccountsRequest
	req.Document_Number = "101"
	_, err := ts.AccountService.CreateAccount(req)
	if err != nil {
		// add logger
	}
}
func TestMain(m *testing.M) {

	setup()
	code := m.Run()
	os.Exit(code)
}
func TestIsValidateRequest(t *testing.T) {

	type data struct {
		inputOperationTypeID int
		inputAccountID       int
		inputAmount          float64
		name                 string
		expected             bool
	}

	tests := []data{data{1, 1, 100, "All valid", true},
		data{0, 1, 100, "InvalidTranTypeID", false},
		data{6, 1, 100, "InvalidTranTypeID", false},
		data{1, 1, -100, "CreateTranAgainWithSameInput", true},
		data{1, 1234, 100, "InvalidAccountID", false},
		data{1, 1, -100, "ValidNegativeAmount", true},
		data{4, 1, -100, "InvalidNegativeAmount", false}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidateRequestSub(test.inputAccountID, test.inputOperationTypeID, test.inputAmount)
			if result != test.expected {
				t.Errorf("%s(%d, %d, %f) = %t; want %t", test.name, test.inputAccountID, test.inputOperationTypeID, test.inputAmount, result, test.expected)
			}
		})

	}

}
