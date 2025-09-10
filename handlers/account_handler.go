package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	customerrors "pismo-code-assessment/customerror"
	"pismo-code-assessment/datastruct"
	"pismo-code-assessment/services"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service *services.AccountService
	Logger  *log.Logger
}

func (a *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()

	a.Logger.Println("Received request to create account")
	a.Logger.Printf("[request_id=%s] handling account creation", requestID) // TODO: further include requestID in all logs to track

	var inputdata datastruct.CreateAccountsRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&inputdata); err != nil {
		a.Logger.Println("Create account request failed!", err)
		errResp := customerrors.ErrorResponse{ErrID: http.StatusBadRequest, Errormsg: "failed to encode input", Details: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	respCreated, err := a.Service.Create(inputdata)
	if err != nil {
		a.Logger.Println("Create account request failed!", err)
		errResp := customerrors.ErrorResponse{ErrID: http.StatusInternalServerError, Errormsg: "failed to create accound", Details: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respCreated)

}

func (a *AccountHandler) GetAccountByIDHandler(w http.ResponseWriter, r *http.Request) {

	requestID := uuid.New().String()

	a.Logger.Println("Received request to get account by ID")
	a.Logger.Printf("[request_id=%s] handling account retrival", requestID) // TODO: further include requestID in all logs to track

	param := mux.Vars(r)
	id := param["accountId"]

	idnumber, _ := strconv.Atoi(id)
	resp, err := a.Service.Get(idnumber)
	if err != nil {
		a.Logger.Println("Get account request failed!", err)
		errResp := customerrors.ErrorResponse{ErrID: http.StatusInternalServerError, Errormsg: "failed to get account", Details: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
