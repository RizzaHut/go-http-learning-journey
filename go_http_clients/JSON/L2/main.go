package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// make an empty slice, with Issue struct as the type, to contain the data 
	var issues []Issue
	
	// decode the JSON data to Go struct
	decoder := json.NewDecoder(res.Body)

	// fill the slice with the decoded data
	if err := decoder.Decode(&issues); err != nil {
		return nil, errors.New("error decoding response body")
	}
	return issues, nil
}
