.PHONY: build clean deploy

download:
	go mod download

build: download
	export GO111MODULE=on
	export GOOS=linux
	go build -ldflags="-s -w" -o bin/healthcheck healthcheck/main.go
	go build -ldflags="-s -w" -o bin/scrape-latest scrape/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
