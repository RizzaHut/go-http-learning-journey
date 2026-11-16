package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// get the data from JSON body, save it to memory, return in byte 
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// make an empty slice, with Issue struct as the type, to contain the data 
	var issues []Issue
	// fill the slice with the byte data 
	if err := json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}
