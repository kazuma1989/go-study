package main

import (
	"log"
)

func main() {
	err := curl("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
}
