package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func walkDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fmt.Println(filepath.Join(dir, file.Name()))

		if file.IsDir() {
			err := walkDir(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
