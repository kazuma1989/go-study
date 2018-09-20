package main

import (
	"encoding/json"
	"fmt"

	"github.com/oliveagle/jsonpath"
)

func jq(path string, input []byte) error {
	var value map[string]interface{}
	err := json.Unmarshal(input, &value)
	if err != nil {
		return err
	}

	filtered, err := jsonpath.JsonPathLookup(value, path)
	if err != nil {
		return err
	}

	b, err := json.Marshal(filtered)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
