. PHONY: fmt vet build

hello:
	@echo "Ready to build some Go! 🐻"

fmt:
	go fmt ./...

vet:
	go vet ./...

build-basics-variables:
	mkdir -pv ./bin/basics
	go build -o ./bin/basics/variables ./basics

run-basics-variables:
	go run ./basics/variables.go

build: build-basics

clean:
	rm -rvf bin