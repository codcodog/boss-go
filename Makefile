target = "bin/boss"

GOOS = "linux"
GOARCH = "amd64"
CGO_ENABLED = 1

.PHONY: all
all: run

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(target)

.PHONY: run
run:
	go run main.go
