package main

import (
	"fmt"
	"net/http"
	"pismo-code-assessment/handlers"
	"pismo-code-assessment/services"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	accountService := services.NewAccountService()
	as := &handlers.AccountHandler{Service: accountService}

	transactionService := services.NewTransactionsService()
	ts := &handlers.TransactionHandler{AccountService: accountService, TransactionService: transactionService}

	router.HandleFunc("/accounts", as.CreateAccountsHandler).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", as.GetAccountsByIDHandler).Methods("GET")
	router.HandleFunc("/transactions", ts.CreateTransactionHandler).Methods("POST")

	return router
}
func main() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", NewRouter())
}
