.PHONY: build run tidy install clean

build:
	go build -o bin/paint cmd/main.go

run:
	go run cmd/main.go

tidy:
	go mod tidy
