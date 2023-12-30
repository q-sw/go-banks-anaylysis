# TO DO list
## Nordigen endpoints
- [x] Create nordigen API endpoint
    - [x] Login [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-1-get-access-token)
    - [x] list all french bank [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-2-choose-a-bank)
    - [x] User agreement [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-3-create-an-end-user-agreement)
    - [x] User requisition [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-4-build-a-link)
    - [x] list account in requisition [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-5-list-accounts)
    - [x] list all transaction from bank account [cf documentation](https://developer.gocardless.com/bank-account-data/quick-start-guide/#step-6-access-account-details-balances-and-transactions)
    - [x] Get bank account balance [cf documentation](https://bankaccountdata.gocardless.com/api/docs#/accounts/retrieve%20account%20balances)
## Bank-analysis endpoints (depends on Storage)
- [ ] Create bank-analysis API endpoint 
    - [ ] categories (get, post, delete)
    - [ ] list all transaction
    - [ ] get bank account balances
    - [ ] get transaction per categories
## Storage
- [ ] Create new connection
- [ ] Create table at connection init
- [ ] Insert data 
- [ ] Update data

## General
- [ ] Reduce the repetitive  
exemple: in api/endpoint.go internal/nordigen.go

- [ ] Improve logging and error handling
    use slog package 
    - [cf. this article](https://gozone.dev/how-to-use-the-new-structured-logging-package-in-go-1-21-d2b49b4f4aed)
    - [slog documentation](https://pkg.go.dev/golang.org/x/exp/slog)
    - [slog go blog article](https://go.dev/blog/slog)
    - [other resources for slog](https://go.dev/wiki/Resources-for-slog)

# Specification 
## Nordigen endpoint login
POST
## Nordigen endpoint banks 
GET
## Nordigen endpoint agree 
POST
## Nordigen endpoint record-bank 
POST
## Nordigen endpoint {record-id}/accounts
GET
Store information in DB 
## Nordigen endpoint {account-id}/transactions
GET
Store latest transactions in DB
## Nordigen endpoint {account-id}/balance
GET
Store balance in DB 

