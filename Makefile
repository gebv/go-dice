build:
	go build -v -o ./bin/dice ./dice/main.go ./dice/collections.go

manual-test: build
	go generate ./examples
