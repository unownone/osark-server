run:
	cd cmd/api && go run *.go

build:
	cd cmd/api && go build -o ../bin/api

generate-docs:
	swag init -g cmd/api/main.go --parseDependency --ot json,yaml

format:
	go fmt ./...

lint:
	golangci-lint run ./...

