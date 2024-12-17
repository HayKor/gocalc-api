.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build -o build/main cmd/server/main.go 
