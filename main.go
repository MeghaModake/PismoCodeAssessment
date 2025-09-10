package main

import (
	"log"
	"net/http"
	"os"
	"pismo-code-assessment/handlers"
	"pismo-code-assessment/services"
	"time"

	"github.com/gorilla/mux"
)

func NewRouter(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	accountService := services.NewAccountService(logger)
	as := &handlers.AccountHandler{Service: accountService, Logger: logger}

	transactionService := services.NewTransactionsService(logger)
	ts := &handlers.TransactionHandler{AccountService: accountService, TransactionService: transactionService, Logger: logger}

	router.Handle("/accounts",
		http.TimeoutHandler(http.HandlerFunc(as.CreateAccountHandler),
			5*time.Second,
			"Request timed out"),
	).Methods("POST")

	router.Handle("/accounts/{accountId}",
		http.TimeoutHandler(http.HandlerFunc(as.GetAccountByIDHandler),
			3*time.Second,
			"Request timed out"),
	).Methods("GET")

	router.Handle("/transactions",
		http.TimeoutHandler(http.HandlerFunc(ts.CreateTransactionHandler),
			5*time.Second,
			"Request timed out"),
	).Methods("POST")

	return router
}
func main() {
	logger := log.New(os.Stdout, "PISMO_LOG: ", log.LstdFlags|log.Lshortfile)
	logger.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", NewRouter(logger)); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
