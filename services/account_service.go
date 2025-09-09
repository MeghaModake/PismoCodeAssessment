package services

import (
	"fmt"
	"log"
	customerrors "pismo-code-assessment/CustomErrors"
	"pismo-code-assessment/datastruct"
	"sync"
)

type AccountService struct {
	mu        *sync.Mutex
	Owners    map[string]int
	Accounts  map[int]string
	LastAccID int
	Logger    *log.Logger
}

func NewAccountService(logger *log.Logger) *AccountService {
	return &AccountService{
		mu:        &sync.Mutex{},
		Owners:    make(map[string]int),
		Accounts:  make(map[int]string),
		LastAccID: 0,
		Logger:    logger,
	}
}

func (as *AccountService) Create(input datastruct.CreateAccountsRequest) (datastruct.Account, error) {

	accid := as.LastAccID + 1
	if _, found := as.Owners[input.Document_Number]; !found {
		as.mu.Lock()
		defer as.mu.Unlock()
		as.Owners[input.Document_Number] = accid
		as.Accounts[accid] = input.Document_Number
		as.LastAccID = accid

		as.Logger.Println("Account created!")
		return datastruct.Account{Account_ID: accid, Document_Number: input.Document_Number}, nil
	} else {
		as.Logger.Printf("Error %s while creating Account with Document_Number %s!\n", customerrors.ACCOUNTEXISTS, input.Document_Number)
		return datastruct.Account{}, fmt.Errorf("%w", customerrors.ACCOUNTEXISTS)
	}

}

func (as *AccountService) Get(id int) (datastruct.Account, error) {

	as.mu.Lock()
	defer as.mu.Unlock()
	if doc, found := as.Accounts[id]; found {
		as.Logger.Println("Account Info retrived !")
		return datastruct.Account{Account_ID: id, Document_Number: doc}, nil
	} else {
		as.Logger.Printf("Get account %d request failed!\n", id)
		return datastruct.Account{}, fmt.Errorf("%w: %d", customerrors.INVALIDACCOUNTID, id)
	}

}
func (as *AccountService) AccountExists(accountid int) bool {

	if _, found := as.Accounts[accountid]; found {
		return true
	}

	return false
}
