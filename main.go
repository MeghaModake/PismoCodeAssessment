package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pismo-code-assessment/handlers"
	"pismo-code-assessment/services"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	logger := log.New(os.Stdout, "APP_LOG: ", log.LstdFlags|log.Lshortfile)

	accountService := services.NewAccountService(logger)
	as := &handlers.AccountHandler{Service: accountService, Logger: logger}

	transactionService := services.NewTransactionsService(logger)
	ts := &handlers.TransactionHandler{AccountService: accountService, TransactionService: transactionService, Logger: logger}

	router.HandleFunc("/accounts", as.CreateAccountsHandler).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", as.GetAccountsByIDHandler).Methods("GET")
	router.HandleFunc("/transactions", ts.CreateTransactionHandler).Methods("POST")

	return router
}
func main() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", NewRouter())
}
