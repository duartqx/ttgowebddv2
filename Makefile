all: dotenv build run

update:
	go get -u ./...

test:
	go test ./... | grep -v '?'

vtest:
	go test -v ./... | grep -v '?'

dotenv:
	set -a && source ./.env

build:
	go build -o ./cmd/main ./cmd

run:
	./cmd/main || true
