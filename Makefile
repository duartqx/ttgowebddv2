test:
	go test ./...

ifeq ($(V),1)
test:
	go test ./... -v
endif
