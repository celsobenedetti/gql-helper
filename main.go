package main

import (
	"encoding/json"
	"log"
	"os"
)

var API = os.Getenv("API")
var X = os.Getenv("X")
var AUTH = os.Getenv("AUTH")
var QUERY_FILE = os.Getenv("QUERY_FILE")
var VARS_FILE = os.Getenv("VARS_FILE")

type JSON = map[string]any

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	query, err := readQueryAndVars()
	if err != nil {
		return err
	}

	client := &Client{
		query: query,
	}

	err = client.doRequest()
	if err != nil {
		return err
	}

	return nil
}

func readQueryAndVars() (*Query, error) {
	query, err := os.ReadFile(QUERY_FILE)
	if err != nil {
		return nil, err
	}

	varsBuf, err := os.ReadFile(VARS_FILE)
	if err != nil {
		return nil, err
	}

	vars := make(JSON)
	err = json.Unmarshal(varsBuf, &vars)
	if err != nil {
		return nil, err
	}

	return &Query{
		Query: string(query),
		Vars:  vars,
	}, nil
}
