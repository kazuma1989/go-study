package main

import (
	"log"
)

func main() {
	err := walkDir("./")
	if err != nil {
		log.Fatal(err)
	}
}
