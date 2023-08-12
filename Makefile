EXEC_FILE := mvc

.PHONY: all replacecredentials&setupMySQL build test run 

all: replacecredentials&setupMySQL build test run open

replacecredentials&setupMySQL:
	chmod +x ./scripts/replaceCredentials.sh
	./scripts/replaceCredentials.sh

build:
	go mod vendor
	go mod tidy
	go build -o $(EXEC_FILE) ./cmd/main.go

test:
	go test ./pkg/models

run:
	./${EXEC_FILE}
