build: lint
	go build -o ./bin/simple ./cmd/simple

lint:
	gofmt -w .

test: lint
	rm -rf mock
	go generate ./...
	go test $$(go list ./... | grep -v mock)

