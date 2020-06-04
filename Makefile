.PHONY:

build:
	go build ./cmd/trading -o bin

run:
	go run ./cmd/trading

pkg-list:
	go list -m all

pkg-clean:
	go clean
	go mod tidy

mod-clean:
	go clean -modcache