# Set the shell to bash always
SHELL := /bin/bash

# Environment variables
ifneq ($(wildcard ./.env),)
	include .env
	export
endif

check: lint test

lint:
	$(LINT) run

tidy:
	go mod tidy

test:
	go test -v ./...
