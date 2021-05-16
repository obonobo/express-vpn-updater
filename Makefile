.PHONY: build clean deploy

download:
	go mod download

build: download
	export GO111MODULE=on
	export GOOS=linux
	go build -ldflags="-s -w" -o bin/hello hello/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose

dev:
	sls local-dev-server --port 3000
