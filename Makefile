.PHONY: setup build cross-build chelp
.DEFAULT_GOAL := help

setup:
	go get github.com/golang/dep/cmd/dep
	go get github.com/Songmu/make2help/cmd/make2help

## Build rafflesia-bot
build: main.go vendor
	go build .

## Build rafflesia-bot for linux
cross-build: main.go vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-w -s" .

## Show help
help:
	@make2help $(MAKEFILE_LIST)

vendor: Gopkg.toml Gopkg.lock
	dep ensure