BINARY_NAME=url-shortener-app-lyalin
.DEFAULT_GOAL := run

build:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${BINARY_NAME}-darwin cmd/url-shortener/main.go

run: build
	./build/${BINARY_NAME}-darwin

test:
	go test ./...

clean:
	go clean
	rm -rf build/