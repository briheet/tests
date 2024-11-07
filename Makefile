run: 
	@go run main.go

test:
	@go test -v ./...

bench:
	@go test -bench=. ./...

gencover: 
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

image:
	@docker build -t gotest:v1 -f Dockerfile.multistage .

runx:
	@docker run -p 5001:5001 gotest:v1
