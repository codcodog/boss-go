target = "bin/boss"

.PHONY: all
all: run

.PHONY: build
build:
	go build -o $(target)

.PHONY: run
run:
	go run main.go
