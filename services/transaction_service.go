package services

import (
	"pismo-code-assessment/datastruct"
	"time"
)

type TransactionsService struct {
	Transactions      map[int]datastruct.Transaction
	LastTransactionID int
}

func NewTransactionsService() *TransactionsService {
	return &TransactionsService{
		Transactions: make(map[int]datastruct.Transaction)}
}

func (ts *TransactionsService) CreateTransaction(req datastruct.CreateTransactionRequest) (datastruct.Transaction, error) {

	trn_id := ts.LastTransactionID + 1
	newTrn := datastruct.Transaction{Transaction_ID: trn_id,
		Account_ID:       req.Account_ID,
		OperationType_ID: req.OperationType_ID,
		Amount:           req.Amount,
		EventDate:        time.Now()}
	ts.Transactions[trn_id] = newTrn

	return newTrn, nil
}
