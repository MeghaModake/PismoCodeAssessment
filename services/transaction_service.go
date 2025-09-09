package services

import (
	"log"
	"pismo-code-assessment/datastruct"
	"sync"
	"time"
)

type TransactionsService struct {
	mu                *sync.Mutex
	Transactions      map[int]datastruct.Transaction
	LastTransactionID int
	Logger            *log.Logger
}

func NewTransactionsService(logger *log.Logger) *TransactionsService {
	return &TransactionsService{
		mu:           &sync.Mutex{},
		Transactions: make(map[int]datastruct.Transaction),
		Logger:       logger}
}

func (ts *TransactionsService) Create(req datastruct.CreateTransactionRequest) (*datastruct.Transaction, error) {
	ts.Logger.Println("Creating Transaction...")
	trnID := ts.LastTransactionID + 1
	newTrn := datastruct.Transaction{Transaction_ID: trnID,
		Account_ID:       req.Account_ID,
		OperationType_ID: req.OperationType_ID,
		Amount:           req.Amount,
		EventDate:        time.Now()}

	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Transactions[trnID] = newTrn
	ts.LastTransactionID = trnID

	ts.Logger.Println("Transaction created!")
	return &newTrn, nil
}
