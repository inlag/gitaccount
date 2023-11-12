all: build

build:
	echo "Build"
	go build -o bin/main ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm ./cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 ./cmd/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 ./cmd/main.go
