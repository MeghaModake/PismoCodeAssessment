# PismoCodeAssessment
Customer Account &amp; Transactions
## Steps:

git clone https://github.com/MeghaModake/PismoCodeAssessment.git
goto location path
go build
go run ./ 

Keep running application and then from another terminal below commands to test the application

// command to test POST /accounts
curl -X POST -H "Content-Type: application/json" -d '{"Document_Number": "101"}'  http://localhost:8080/accounts

// command to test GET /accounts/id
curl -X GET "localhost:8080/accounts/1" -H "Content-Type: application/json"
Can also use browser to see response 

// command to test POST /transactions
curl -X POST -H "Content-Type: application/json" -d '{"account_id": 1, "operation_type_id":4, "amount": 123.45}'  http://localhost:8080/transactions