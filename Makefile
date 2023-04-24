.DEFAULT_GOAL:=help
SHELL:=/bin/sh

APP_NAME=alina--tool


.PHONY: build
build:
	echo "Building $(APP_NAME)"
	go build ./cmd/api/