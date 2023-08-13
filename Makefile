EXEC_FILE := mvc

.PHONY: all configFileBuild&setupMySQL setup build test run 

all: configFileBuild&setupMySQL setup build test run 

configFileBuild&setupMySQL:
	chmod +x ./scripts/configFileBuild.sh
	./scripts/configFileBuild.sh

setup:
	python3 ./scripts/setup.py

build:
	go mod vendor
	go mod tidy
	go build -o $(EXEC_FILE) ./cmd/main.go

test:
	go test ./pkg/models

run:
	./$(EXEC_FILE)

