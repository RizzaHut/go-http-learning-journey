package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var marshalData [][]byte

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		marshalData = append(marshalData, data)
	}

	return marshalData, nil
}
