all: build run

update:
	go get -u ./...

test:
	cd ./internal && go test ./... | grep -v '?'; cd ./..

vtest:
	cd ./internal && go test ./... -v | grep -v '?'; cd ./..

dotenv:
	set -a && source ./.env

build:
	go build -o ./cmd/main ./cmd

run:
	./cmd/main || true
