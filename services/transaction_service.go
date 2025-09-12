package services

import (
	"log"
	"pismo-code-assessment/datastruct"
	"sort"
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

/*
type Transaction struct {
	Transaction_ID   int
	Account_ID       int
	OperationType_ID int
	Amount           float64
	EventDate        time.Time
}
*/

func (ts *TransactionsService) ListNegativeTransactions(accid int) []datastruct.Transaction {

	ts.Logger.Println("Listing old Transactions!")

	listtransaction := make([]datastruct.Transaction, 0, 0)
	for _, v := range ts.Transactions {
		if v.Account_ID == accid && v.Amount < 0 && v.Balance < 0 {
			listtransaction = append(listtransaction, v)
			ts.Logger.Println("Added Transaction into list!", v.Transaction_ID)
		}
	}

	return listtransaction
}

func (ts *TransactionsService) UpdateEarlierTransactionBalance(req datastruct.CreateTransactionRequest, oldTrns []datastruct.Transaction, amount float64) float64 {

	ts.Logger.Println("Updating old Transactions!", len(oldTrns))

	//Extract values into a slice
	transactions := make([]datastruct.Transaction, 0, len(oldTrns))
	for _, tx := range oldTrns {
		transactions = append(transactions, tx)
	}

	//Sort by EventDate (earliest first)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].EventDate.Before(transactions[j].EventDate)
	})

	remainingbalance := amount
	for _, v := range transactions {
		var currbalance float64
		remainingbalance = v.Balance + remainingbalance
		if remainingbalance >= 0 {
			currbalance = 0
		} else {
			currbalance = remainingbalance
		}

		ts.mu.Lock()
		newTrn := datastruct.Transaction{Transaction_ID: v.Transaction_ID,
			Account_ID:       v.Account_ID,
			OperationType_ID: v.OperationType_ID,
			Amount:           v.Amount,
			Balance:          currbalance,
			EventDate:        v.EventDate}

		ts.Transactions[v.Transaction_ID] = newTrn
		ts.mu.Unlock()

		if remainingbalance <= 0 {
			remainingbalance = 0
			break
		}
	}

	return remainingbalance

}

/*
1 1 -50 -50
1 1 -40 -40
4 1 -90


1 1 -50 -50	0
1 1 -40 -40	0
4 1 -90

//
1 1 -50 -50
1 1 -40 -40
4 1 60


1 1 -50 -50	0
1 1 -40 -40	-30
4 1  60

*/

func (ts *TransactionsService) Create(req datastruct.CreateTransactionRequest) (*datastruct.Transaction, error) {
	ts.Logger.Println("Creating Transaction...")

	//Get earliers transactions map where account ID = id ,amount < 0 && balance < 0

	var newTrn datastruct.Transaction
	if req.OperationType_ID == datastruct.OpPayment {

		ts.Logger.Println("List and update old Transactions!")

		listtrns := ts.ListNegativeTransactions(req.Account_ID)
		newbalance := ts.UpdateEarlierTransactionBalance(req, listtrns, req.Amount)

		ts.Logger.Println("here!")
		trnID := ts.LastTransactionID + 1
		newTrn = datastruct.Transaction{Transaction_ID: trnID,
			Account_ID:       req.Account_ID,
			OperationType_ID: req.OperationType_ID,
			Amount:           req.Amount,
			Balance:          newbalance,
			EventDate:        time.Now()}

		ts.mu.Lock()
		defer ts.mu.Unlock()
		ts.Transactions[trnID] = newTrn
		ts.LastTransactionID = trnID

	} else {

		ts.Logger.Println("Just added a new Transaction!")

		trnID := ts.LastTransactionID + 1
		newTrn = datastruct.Transaction{Transaction_ID: trnID,
			Account_ID:       req.Account_ID,
			OperationType_ID: req.OperationType_ID,
			Amount:           req.Amount,
			Balance:          req.Amount,
			EventDate:        time.Now()}

		ts.mu.Lock()
		defer ts.mu.Unlock()
		ts.Transactions[trnID] = newTrn
		ts.LastTransactionID = trnID

	}

	//Extract values into a slice
	transactions := make([]datastruct.Transaction, 0, len(ts.Transactions))
	for _, tx := range ts.Transactions {
		transactions = append(transactions, tx)
	}

	//Sort by EventDate (earliest first)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].EventDate.Before(transactions[j].EventDate)
	})

	ts.Logger.Println("Printing all transactions!")
	for _, v := range ts.Transactions {
		ts.Logger.Println(v)
	}

	ts.Logger.Println("Transaction created!")

	return &newTrn, nil
}
