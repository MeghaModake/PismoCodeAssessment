package datastruct

import "time"

type CreateTransactionRequest struct {
	OperationType_ID int     `json:"operation_type_id"`
	Account_ID       int     `json:"account_id"`
	Amount           float64 `json:"amount"`
}

type Transaction struct {
	Transaction_ID   int
	Account_ID       int
	OperationType_ID int
	Amount           float64
	EventDate        time.Time
}

// Not used currently but just kept it here for reference
var operations = map[int]string{
	1: "PURCHASE",
	2: "INSTALLMENT PURCHASE",
	3: "WITHDRAWAL",
	4: "PAYMENT",
}
