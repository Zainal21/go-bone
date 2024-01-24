BINARY_NAME=go-bone-boilerplate
build:
	@go build -o bin/${BINARY_NAME} main.go

run-http:
	@./bin/${BINARY_NAME} http
	
install:
	@echo "Installing dependencies...."
	@rm -rf vendor
	@rm -f Gopkg.lock
	@rm -f glide.lock
	@go mod tidy && go mod download && go mod vendor

start-http:
	@go run main.go http


start-rabbit-mq:
	@go run main.go rabbit

	
start-rabbit-mq:
	@go run main.go rabbit

migrate:
	@go run main.go db:migrate