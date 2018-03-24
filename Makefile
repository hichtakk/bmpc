.PHONEY: test
test:
	go test -cover -v

.PHONEY: build
build:
	go build
