swagger init:
	swag init -g main.go --output docs
build:
	go build -o . /dummy_service
run:
	go run .