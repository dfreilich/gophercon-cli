GOCMD?=go
NAME?=mycli

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

.PHONY: build test