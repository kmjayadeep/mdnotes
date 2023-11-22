all: build test

build:
	go build -o bin/mdnotes cmd/mdnotes/main.go

test:
	go test -v ./...

css:
	@tailwindcss -i ./assets/app.css -o ./assets/dist/app.css 

run: 
	@go run cmd/mdnotes/main.go

fmt:
	@go fmt ./...

clean:
	go clean
	rm ./bin/mdnotes
