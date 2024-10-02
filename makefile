


test:
	go test -v -race ./... 
coverage: 
	go test -v -race -cover -coverprofile=.cover.out ./...
	go tool cover -html=.cover.out
lint:
	golangci-lint run ./...
