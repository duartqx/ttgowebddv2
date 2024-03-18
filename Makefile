test:
	go test ./... | grep -v '?'

ifeq ($(V),1)
test:
	go test ./... -v | grep -v '?'
endif
