package services

import (
	"fmt"
	"pismo-code-assessment/datastruct"
)

type AccountService struct {
	owners    map[string]int
	accounts  map[int]string
	lastAccID int
}

func NewAccountService() *AccountService {
	return &AccountService{
		owners:   make(map[string]int),
		accounts: make(map[int]string),
	}
}

func (as *AccountService) CreateAccount(input datastruct.CreateAccountsRequest) (datastruct.Account, error) {

	accid := as.lastAccID + 1
	if _, found := as.owners[input.Document_Number]; !found {
		as.owners[input.Document_Number] = accid
		as.accounts[accid] = input.Document_Number
		as.lastAccID = accid
		return datastruct.Account{Account_ID: accid, Document_Number: input.Document_Number}, nil
	} else {
		return datastruct.Account{}, fmt.Errorf("Account Already exists")
	}

}

func (as *AccountService) GetAccount(id int) (datastruct.Account, error) {

	if doc, found := as.accounts[id]; found {
		return datastruct.Account{Account_ID: id, Document_Number: doc}, nil
	} else {
		return datastruct.Account{}, fmt.Errorf("Account does not exists")
	}

}
