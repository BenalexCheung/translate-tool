#!/bin/bash

APP_PATH=build

all: build

build:
	go build .

run:
	bee run

pack: build
	mkdir -p $(APP_PATH)
	cp translateTool $(APP_PATH)
	cp -fr views $(APP_PATH)
	cp -fr static $(APP_PATH)
	cp -fr conf $(APP_PATH)
	cp -fr db $(APP_PATH)
	cp Dockerfile $(APP_PATH)
	cd $(APP_PATH)
	docker build -t pantum/translate-tool:v0.1 .

clean:
	rm -rf ./build translateTool
