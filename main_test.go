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
)

func TestCreateAndGetAccount(t *testing.T) {
	logger := log.New(os.Stdout, "TEST_log", log.LstdFlags)
	router := NewRouter(logger)

	// Create account
	acc := datastruct.CreateAccountsRequest{Document_Number: "100"}
	body, _ := json.Marshal(acc)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 400, got %d", w.Code)
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

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.Document_Number != "101" || got.Account_ID != 2 {
		t.Errorf("unexpected account: %+v", got)
	}
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
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

	req = httptest.NewRequest("GET", "/accounts/3", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}

}

func TestCreateAndGetTransaction(t *testing.T) {
	logger := log.New(os.Stdout, "TEST_log", log.LstdFlags)
	router := NewRouter(logger)

	tx := datastruct.CreateTransactionRequest{Account_ID: 1, OperationType_ID: 1, Amount: 50}
	body, _ := json.Marshal(tx)

	// Create transaction
	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got datastruct.Transaction
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Transaction_ID != 1 || got.Account_ID != 1 {
		t.Errorf("unexpected ID: %+v", got)
	}

}
