package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func prettyJSON(b []byte) {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, b, "", "  ")
	fmt.Println(string(prettyJSON.Bytes()))
}

func prettyStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func printResponse(res *http.Response) error {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	prettyJSON(resBody)
	return nil
}
