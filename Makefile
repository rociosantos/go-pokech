.PHONY: unit-test
unit-test:
	@echo "::: running unit tests"
	go test -v ./...

.PHONY: lint
lint:
	@echo "::: running code lint"
	golangci-lint run ./... --config=.golangci.yml

.PHONY: deps
deps:
	@echo "::: installing golang dependencies"
	go mod tidy

.PHONY: mocks
mocks:
	@echo "::: generating mocks"
	go generate -x ./...

.PHONY: run
run:
	go run main.go
