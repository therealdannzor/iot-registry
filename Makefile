.DEFAULT_GOAL := build

BINARY_NAME = start

build:
	@go build -o $(BINARY_NAME)

clean:
	@rm -f $(BINARY_NAME) 
