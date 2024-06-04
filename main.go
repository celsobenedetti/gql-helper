package main

import (
	"fmt"
	"os"
)

var TOKEN = os.Getenv("TOKEN")
var QUERY_FILE = os.Getenv("QUERY_FILE")
var VARS_FILE = os.Getenv("VARS_FILE")

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	query, err := readQueryAndVars()
	if err != nil {
		return err
	}

	fmt.Println(query)

	return nil
}

func readQueryAndVars() (*Query, error) {
	query, err := os.ReadFile(QUERY_FILE)
	if err != nil {
		return nil, err
	}

	vars, err := os.ReadFile(VARS_FILE)
	if err != nil {
		return nil, err
	}

	return &Query{
		query: string(query),
		vars:  string(vars),
	}, nil
}

type Query struct {
	query string
	vars  string
}
