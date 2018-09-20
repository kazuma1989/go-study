package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	l := len(args)
	var url string
	switch {
	case l == 0:
		log.Fatal("Need 1 argument.")
	case l == 1:
		url = args[0]
	case l >= 2:
		log.Fatal("Too many arguments.")
	}

	err := curl(url, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
