GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
VETPACKAGES ?= $(shell $(GO) list ./...)
CLI_VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
GO_LDFLAGS ?= -X $(shell $(GO) list -m)/cmd.Version=$(CLI_VERSION)
VER ?= $(shell git describe --tags --abbrev=0)
TAGS ?=

bina_json = '{"platforms": { "darwin-arm64": { "asset": "snake-${VER}-arm64-Darwin.tar.gz", "file": "snake" }, "darwin-amd64": { "asset": "snake-${VER}-x86_64-Darwin.tar.gz", "file": "snake" }, "linux-arm64": { "asset": "snake-${VER}-arm64-Linux.tar.gz", "file": "snake" }, "linux-amd64": { "asset": "snake-${VER}-x86_64-Linux.tar.gz", "file": "snake" }, "windows-amd64": { "asset": "snake-${VER}-x86_64-Windows.tar.gz", "file": "snake.exe" } } }'

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

vet:
	$(GO) vet $(VETPACKAGES)

build:
	$(GO) build -tags "$(TAGS)" -o bin/snake -ldflags "-s -w ${GO_LDFLAGS}" *.go

archive-release:
	rm -rf bin/snake
	GOARCH=arm64 GOOS=darwin $(GO) build -tags "$(TAGS)" -o bin/snake -ldflags "-s -w ${GO_LDFLAGS}" *.go
	tar -C ./bin -czf bin/snake-${VER}-arm64-Darwin.tar.gz snake
	rm -rf bin/snake
	GOARCH=amd64 GOOS=darwin $(GO) build -tags "$(TAGS)" -o bin/snake -ldflags "-s -w ${GO_LDFLAGS}" *.go
	tar -C ./bin -czf bin/snake-${VER}-x86_64-Darwin.tar.gz snake
	rm -rf bin/snake
	GOARCH=arm64 GOOS=linux $(GO) build -tags "$(TAGS)" -o bin/snake -ldflags "-s -w ${GO_LDFLAGS}" *.go
	tar -C ./bin -czf bin/snake-${VER}-arm64-Linux.tar.gz snake
	rm -rf bin/snake
	GOARCH=amd64 GOOS=linux $(GO) build -tags "$(TAGS)" -o bin/snake -ldflags "-s -w ${GO_LDFLAGS}" *.go
	tar -C ./bin -czf bin/snake-${VER}-x86_64-Linux.tar.gz snake
	rm -rf bin/snake
	make bina

bina:
	@echo ${bina_json} > ./bin/bina.json

version:
	@echo ${VER} > ./VERSION
