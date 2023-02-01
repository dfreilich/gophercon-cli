GOCMD?=go
NAME?=joker

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

.PHONY: build test