package main

import (
	"encoding/json"
	"fmt"
)

func jq(path string, input []byte) error {
	var value map[string]string
	err := json.Unmarshal(input, &value)
	if err != nil {
		return err
	}

	fmt.Println(value)
	return nil
}
