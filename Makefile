SHELL := /bin/bash

APP_NAME = wasm_builder


build: ## build wasm
	GOOS=js GOARCH=wasm go build -o live/wok.wasm wasm/wok.go
