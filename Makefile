GOCMD?=go
NAME?=joker

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

run:
	$(GOCMD) run ./...

.PHONY: build test run