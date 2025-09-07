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
	ai := &handlers.AccountHandler{Service: accountService}

	router.HandleFunc("/accounts", ai.CreateAccountsHandler).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", ai.GetAccountsByIDHandler).Methods("GET")

	return router
}
func main() {

	http.ListenAndServe(":8080", NewRouter())
	fmt.Println("Server is running on port 8080")

}
