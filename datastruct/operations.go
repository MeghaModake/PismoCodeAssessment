package datastruct

import "time"

type operation_types struct {
	OperationType_ID int    `json:"operation_type_id"`
	Description      string `json:"description"`
}

var operations = map[int]string{
	1: "PURCHASE",
	2: "INSTALLMENT PURCHASE",
	3: "WITHDRAWAL",
	4: "PAYMENT",
}

type CreateTransactionRequest struct {
	OperationType_ID int `json:"operation_type_id"`
	Account_ID       int `json:"account_id"`
	Amount           int `json:"amount"`
}

type Transaction struct {
	Transaction_ID   int
	Account_ID       int
	OperationType_ID int
	Amount           int
	EventDate        time.Time
}
