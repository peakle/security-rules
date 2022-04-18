GOPATH_DIR=$(shell go env GOPATH)
VERSION=$(shell git describe --tags)

test:
	go test --count=1 -race .

lint:
	golangci-lint run ./...
	golangci-lint run ./testdata/src/...
