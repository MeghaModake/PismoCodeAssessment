package main

import (
	"log"
	"net/http"
	"os"
	"pismo-code-assessment/handlers"
	"pismo-code-assessment/services"

	"github.com/gorilla/mux"
)

func NewRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

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
	logger := log.New(os.Stdout, "PISMO_LOG: ", log.LstdFlags|log.Lshortfile)
	logger.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", NewRouter(logger)); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
