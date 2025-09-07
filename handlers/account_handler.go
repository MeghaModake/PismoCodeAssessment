package handlers

import (
	"encoding/json"
	"net/http"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service *services.AccountService
}

func (a *AccountHandler) CreateAccountsHandler(w http.ResponseWriter, r *http.Request) {

	var inputdata datastruct.CreateAccountsRequest
	if err := json.NewDecoder(r.Body).Decode(&inputdata); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respCreated, err := a.Service.CreateAccount(inputdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(respCreated)

}

func (a *AccountHandler) GetAccountsByIDHandler(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id := param["accountId"]

	idnumber, _ := strconv.Atoi(id)
	resp, err := a.Service.GetAccount(idnumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
