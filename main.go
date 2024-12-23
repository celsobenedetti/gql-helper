package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var API = os.Getenv("API")
var X = os.Getenv("X")
var AUTH = os.Getenv("AUTH")
var QUERY_FILE = os.Getenv("QUERY_FILE")
var VARS_FILE = os.Getenv("VARS_FILE")

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

func (c *Client) doRequest() error {
	body, err := json.Marshal(c.query)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, API, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("X-Contextual-Entity", X)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AUTH)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if err = c.prettyResponse(res); err != nil {
		return err
	}

	return nil
}

func (c *Client) prettyResponse(res *http.Response) error {
	fmt.Println(res.Status)
	fmt.Println(c.query.Vars)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, resBody, "", "  ")
	fmt.Println(string(prettyJSON.Bytes()))
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

type JSON = map[string]any

type Client struct {
	query *Query
}

type Query struct {
	Query string `json:"query"`
	Vars  JSON   `json:"variables"`
}
