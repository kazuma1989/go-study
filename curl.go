package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func curl(url string, input io.Reader) error {
	resp, err := http.Post(url, "application/json", input)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
