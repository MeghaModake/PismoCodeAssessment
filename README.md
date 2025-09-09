## Project Setup & Execution Guide

## Prerequisites

1.Install Docker Desktop
Download docker desktop https://docs.docker.com/desktop/release-notes/ 
Open and run docker desktop

## Setup Instructions
git clone https://github.com/MeghaModake/PismoCodeAssessment.git

goto location path

go build

docker build -t pismo-assignment-megha .


## Running & Testing the Project

docker run -p 8080:8080 pismo-assignment-megha

// command to test POST /accounts
curl -X POST -H "Content-Type: application/json" -d '{"Document_Number": "101"}'  http://localhost:8080/accounts

// command to test GET /accounts/id
curl -X GET "localhost:8080/accounts/1" -H "Content-Type: application/json"

// command to test POST /transactions
curl -X POST -H "Content-Type: application/json" -d '{"account_id": 1, "operation_type_id":4, "amount": 123.45}'  http://localhost:8080/transactions


## Running & Testing Swagger documentation

docker run -p 8081:8080 \
  -e SWAGGER_JSON=/pismoapp/openapi.yaml \
  -v $(pwd):/pismoapp \
  swaggerapi/swagger-ui
  
 Open http://localhost:8081/