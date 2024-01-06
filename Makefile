.PHONY: test

dev:
	@gochange -k -i '**/*.go' -- make test

test:
	@go test ./...

.DEFAULT_GOAL := dev 