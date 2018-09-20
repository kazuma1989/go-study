package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	l := len(args)
	var path string
	switch {
	case l == 0:
		log.Fatal("Need 1 argument.")
	case l == 1:
		path = args[0]
	case l >= 2:
		log.Fatal("Too many arguments.")
	}

	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		log.Fatal(err)
	} else {
		err := jq(path, input)
		if err != nil {
			log.Fatal(err)
		}
	}
}
