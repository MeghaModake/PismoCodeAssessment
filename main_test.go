package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pismo-code-assessment/datastruct"
	"testing"
)

func TestCreateAndGetAccount(t *testing.T) {
	router := NewRouter()

	// Create account
	acc := datastruct.CreateAccountsRequest{Document_Number: "100"}
	body, _ := json.Marshal(acc)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var got datastruct.CreateAccountsResponse
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

	// // Create another account with non unique Document_Number
	acc = datastruct.CreateAccountsRequest{Document_Number: "101"}
	body, _ = json.Marshal(acc)

	req = httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get account
	req = httptest.NewRequest("GET", "/accounts/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Document_Number != "Alice" || got.Account_ID != 1 {
		t.Errorf("unexpected account: %+v", got)
	}
}

/*func TestCreateAndGetTransaction(t *testing.T) {
	router := NewRouter()

	tx := models.Transaction{ID: "t1", AccountID: "1", Amount: 50}
	body, _ := json.Marshal(tx)

	// Create transaction
	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get transaction
	req = httptest.NewRequest("GET", "/transactions/t1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var got models.Transaction
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	if got.Amount != 50 {
		t.Errorf("unexpected transaction: %+v", got)
	}
}
*/
