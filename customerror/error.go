package customerrors

import "fmt"

var ACCOUNTEXISTS = fmt.Errorf("Account Already exists")
var INVALIDACCOUNTID = fmt.Errorf("Wrong Account ID, Account does not exists")

type ErrorResponse struct {
	ErrID    int    `json:"error_id"`
	Errormsg string `json:"error_message"`
	Details  string `json:"error_details"`
}
