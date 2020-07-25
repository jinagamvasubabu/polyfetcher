GO111MODULE=on

build:
	export GO111MODULE on; \
	go build ./...

lint: build
	golint -set_exit_status .

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out -short

test: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out
