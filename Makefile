run: 
	@go run main.go

test:
	@go test -v ./...

bench:
	@go test -bench=. ./...

gencover: 
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
