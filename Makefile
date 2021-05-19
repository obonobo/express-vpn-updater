.PHONY: build clean deploy
SHELL = bash
STAGE ?= dev

download:
	go mod download

build: download
	export GO111MODULE=on
	export GOOS=linux
	go build -ldflags="-s -w" -o bin/healthcheck   healthcheck/handler.go
	go build -ldflags="-s -w" -o bin/scrape-latest scrape/handlers/get-latest-link/handler.go
	go build -ldflags="-s -w" -o bin/scratch scratch/handler.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose --stage "${STAGE}"
