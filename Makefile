GO111MODULE=on

test: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out
