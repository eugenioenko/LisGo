SHELL := /bin/bash

APP_NAME = wasm_builder

build:
	GOOS=js GOARCH=wasm go build -o live/lisgo.wasm wasm/lisgo.go
test:
	go test ./tests -v
