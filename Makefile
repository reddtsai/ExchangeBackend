.PHONY:

build:
	go build -o bin ./cmd/trading

run:
	go run ./cmd/trading

pkg-list:
	go list -m all

pkg-list-u:
	go list -m -u all

pkg-get:
ifdef package
	# go get -u github.com/reddtsai/go-ringing@v1.0.0
	go get -u $(package)
else
	@echo 'make pkg-get package=PackageName'
endif

pkg-clean:
	go clean
	go mod tidy

mod-clean:
	go clean -modcache

docker-build:
	docker build -t exchange-trading -f build/trading/Dockerfile .
	docker build -t exchange-notification -f build/notification/Dockerfile .

docker-run:
	docker-compose -f build/docker-compose.yml up -d 
