watch-dev:
	air

dev:
	go run cmd/main.go

build:
	go build -o bin/app main.go