package main

import (
	"log"
)

func main() {
	err := WalkDir("./")
	if err != nil {
		log.Fatal(err)
	}
}
