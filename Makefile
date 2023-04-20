.PHONY: build
build:
	go build -v  -o "./bin/server" ./cmd/server

.PHONY: test
test:
	go test -race -v -cover -coverprofile=coverage_usecase.out -timeout 40s ./internal/usecase/shortener
	go test -race -v -cover -coverprofile=coverage_repo.out -timeout 40s ./internal/repository/memory

.PHONY: test-report
test-report:
	make test
	go tool cover -html=coverage_usecase.out
	go tool cover -html=coverage_repo.out
	rm *.out

.DEFAULT_GOAL := build