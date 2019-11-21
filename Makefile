all: unit

unit:
	go test -tags=unit ./...
	go test -tags=unit ./... -race
	go vet -tags=unit ./...
.PHONY: unit