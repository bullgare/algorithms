LAST_GOPATH_DIR:=$(lastword $(subst :, ,$(GOPATH)))
GOBIN:=$(LAST_GOPATH_DIR)/bin
BINPATH:=$(GOBIN)/bualgo

build:
	go build -o $(BINPATH) ./cmd/main.go

test:
	go test ./...

.PHONY: build test