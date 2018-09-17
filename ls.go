package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// WalkDir はディレクトリの中身を再帰的にリストアップします。
func WalkDir(dir string) error {
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
			err := WalkDir(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
