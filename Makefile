target = "bin/boss"

GOOS = "linux"
ARCH = "amd64"
CGO_ENABLED = 0

.PHONY: all
all: run

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) ARCH=$(ARCH) go build -o $(target)

.PHONY: run
run:
	go run main.go
