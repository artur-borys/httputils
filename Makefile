GOBIN ?= $$(go env GOPATH)/bin

.PHONY: test test-coverage install-go-test-coverage check-coverage coverage-html
test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...

install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

check-coverage: install-go-test-coverage test-coverage
	${GOBIN}/go-test-coverage --config=./.testcoverage.yaml

coverage-html: test-coverage
	go tool cover -html=cover.out -o=cover.html
