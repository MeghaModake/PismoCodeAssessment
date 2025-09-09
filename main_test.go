package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"pismo-code-assessment/datastruct"
	"testing"

	"github.com/gorilla/mux"
)

var router *mux.Router
var logger *log.Logger

func setup() {
	logger = log.New(os.Stdout, "TEST_log", log.LstdFlags)
	router = NewRouter(logger)
}
func TestMain(m *testing.M) {

	setup()
	code := m.Run()
	os.Exit(code)
}

func TestCreateAndGetAccount(t *testing.T) {

	// Create account
	acc := datastruct.CreateAccountsRequest{Document_Number: "100"}
	body, _ := json.Marshal(acc)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
	var got datastruct.Account
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.Document_Number != "100" || got.Account_ID != 1 {
		t.Errorf("unexpected account: %+v", got)
	}
	// // Create another account
	acc = datastruct.CreateAccountsRequest{Document_Number: "101"}
	body, _ = json.Marshal(acc)

	req = httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.Document_Number != "101" || got.Account_ID != 2 {
		t.Errorf("unexpected account: %+v", got)
	}
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
	// // Create another account with non unique Document_Number
	acc = datastruct.CreateAccountsRequest{Document_Number: "101"}
	body, _ = json.Marshal(acc)

	req = httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Get account 1
	req = httptest.NewRequest("GET", "/accounts/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Document_Number != "100" || got.Account_ID != 1 {
		t.Errorf("unexpected account: %+v", got)
	}

	// Get account 2
	req = httptest.NewRequest("GET", "/accounts/2", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Document_Number != "101" || got.Account_ID != 2 {
		t.Errorf("unexpected account: %+v", got)
	}

	// negative test , should fail
	req = httptest.NewRequest("GET", "/accounts/3", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

}

func TestCreateAndGetTransaction(t *testing.T) {

	// Create account first
	acc := datastruct.CreateAccountsRequest{Document_Number: "100"}
	body, _ := json.Marshal(acc)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// create transaction
	tx := datastruct.CreateTransactionRequest{Account_ID: 1, OperationType_ID: 1, Amount: 50}
	body, _ = json.Marshal(tx)

	// Create transaction
	req = httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
	var got datastruct.Transaction
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Transaction_ID != 1 || got.Account_ID != 1 {
		t.Errorf("unexpected ID: %+v", got)
	}

}
