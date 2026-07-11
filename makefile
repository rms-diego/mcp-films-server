dev-watch:
	air

dev:
	go run cmd/main.go

build-app:
	go build -o build/app cmd/main.go

run-app:
	./build/app