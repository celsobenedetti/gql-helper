package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	query *Query
}

type Query struct {
	Query string `json:"query"`
	Vars  JSON   `json:"variables"`
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
