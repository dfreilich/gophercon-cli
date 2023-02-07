GOCMD?=go
NAME?=joker

build:
	$(GOCMD) build -o $(NAME)

test:
	$(GOCMD) test ./... -v

run:
	$(GOCMD) run ./...

clean:
	git reset . && git checkout . && git clean -ffd

.PHONY: build clean test run