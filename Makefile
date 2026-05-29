.PHONY: build run tidy install clean

build:
	go build -o bin/paint main/main.go

run:
	go run main/main.go

tidy:
	go mod tidy
