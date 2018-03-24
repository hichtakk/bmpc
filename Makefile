.PHONY: test
test:
	go test -cover -v

.PHONY: build
build:
	go build
