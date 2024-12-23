.PHONY: test clean all

build:
	go build -o ~/.local/bin/gql

run: build
	~/.local/bin/gql
