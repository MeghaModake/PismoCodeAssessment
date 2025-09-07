package datastruct

type Account struct {
	Account_ID      int    `json:"account_id"`
	Document_Number string `json:"document_number"`
}

type CreateAccountsRequest struct {
	Document_Number string `json:"document_number"`
}
