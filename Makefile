all: build run

pull-base:
	git remote update && git pull base base

update:
	go get -u ./...

test:
	cd ./internal && go test ./... | grep -v '?'; cd ./..

vtest:
	cd ./internal && go test ./... -v | grep -v '?'; cd ./..

dotenv:
	set -a && source ./.env

build:
	go build -o main .

run:
	./main || true
