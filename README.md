# LMS_GO

## About 
This is the repo for Library Management System written in go following MVC architecture.
###Features
1. Separate admin and client portal with required special privileges for admin.\
2. JWT Authentication and Salting & Hashing for storing passwords.
3. User friendly and attractive interface.

## Application Flow
###Admin-Side
1. Admin can view the books inventory and add, delete or remove the books.
2. He can see the issue and return requests and can accept or reject them.
3. He can add and remove admin.

###User-Side
1. User can view the books available in library and also the books issued against their account.
2. He can accept/withdraw issue of books.
3. He can accept/withdraw return of books.

## Manual Setup
Run the following commands on your terminal
```
git clone https://github.com/Altoiids/LMS_GO.git
cd LMS_GO
go mod vendor
go mod tidy
```
### MySQL:
Run below commands after entering your respective credentials
1. Run this command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" -verbose up`
2. If this command gives error run command : `migrate -path database/migration/ -database "mysql://your_db_username:your_db_password@tcp(localhost:3306)/DB_NAME" force version`


### Run the test function:
1. run command `go test ./pkg/models` .
2. Ensure that test passes.

### Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

### Accessing the website
1. Open localhost:8000 on your browser
## Setup using Makefile
Run the make command in your root directory
