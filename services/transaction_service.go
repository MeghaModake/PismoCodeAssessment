package services

import (
	"log"
	"pismo-code-assessment/datastruct"
	"time"
)

type TransactionsService struct {
	Transactions      map[int]datastruct.Transaction
	LastTransactionID int
	Logger            *log.Logger
}

func NewTransactionsService(logger *log.Logger) *TransactionsService {
	return &TransactionsService{
		Transactions: make(map[int]datastruct.Transaction),
		Logger:       logger}
}

func (ts *TransactionsService) CreateTransaction(req datastruct.CreateTransactionRequest) (datastruct.Transaction, error) {
	ts.Logger.Println("Creating Transaction...")
	trn_id := ts.LastTransactionID + 1
	newTrn := datastruct.Transaction{Transaction_ID: trn_id,
		Account_ID:       req.Account_ID,
		OperationType_ID: req.OperationType_ID,
		Amount:           req.Amount,
		EventDate:        time.Now()}
	ts.Transactions[trn_id] = newTrn
	ts.LastTransactionID = trn_id
	ts.Logger.Println("Transaction created!")
	return newTrn, nil
}
