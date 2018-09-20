package main

import (
	"encoding/json"
	"fmt"
)

func jq(path string, input []byte) error {
	var value map[string]interface{}
	err := json.Unmarshal(input, &value)
	if err != nil {
		return err
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
