package services

import (
	"fmt"
	"log"
	"pismo-code-assessment/datastruct"
)

type AccountService struct {
	Owners    map[string]int
	Accounts  map[int]string
	LastAccID int
	Logger    *log.Logger
}

func NewAccountService(logger *log.Logger) *AccountService {
	return &AccountService{
		Owners:    make(map[string]int),
		Accounts:  make(map[int]string),
		LastAccID: 0,
		Logger:    logger,
	}
}

func (as *AccountService) CreateAccount(input datastruct.CreateAccountsRequest) (datastruct.Account, error) {

	accid := as.LastAccID + 1
	if _, found := as.Owners[input.Document_Number]; !found {
		as.Owners[input.Document_Number] = accid
		as.Accounts[accid] = input.Document_Number
		as.LastAccID = accid
		as.Logger.Println("Account created!")
		return datastruct.Account{Account_ID: accid, Document_Number: input.Document_Number}, nil
	} else {
		return datastruct.Account{}, fmt.Errorf("Account Already exists")
	}

}

func (as *AccountService) GetAccount(id int) (datastruct.Account, error) {

	if doc, found := as.Accounts[id]; found {
		return datastruct.Account{Account_ID: id, Document_Number: doc}, nil
	} else {
		return datastruct.Account{}, fmt.Errorf("Account does not exists")
	}

}
func (as *AccountService) AccountExits(accountid int) bool {

	if _, found := as.Accounts[accountid]; found {
		return true
	}

	return false
}
