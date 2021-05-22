SHELL := /bin/bash
DATE := $(shell date +%Y%m%d-%H:%M:%S)
HASH := $(shell git rev-parse HEAD)

godog:
	go version
	SQL_MODE=debug go test -v --godog.format=pretty
# 	go test -v --godog.format=progress -coverprofile=coverage.out -covermode=atomic
# 	go tool cover -html=coverage.out -o coverage.html
