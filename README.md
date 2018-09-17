# Go に触れてみよう

## インストール

https://golang.org/doc/install#install


## CLI ツールを作る

### `ls`

#### 再帰しない版

`main.go`

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```

`go run main.go` で実行。


#### 再帰版

```diff
+ func walkDir(dir string) error {
+ 	files, err := ioutil.ReadDir(dir)
+ 	if err != nil {
+ 		return err
+ 	}
+ 
+ 	for _, file := range files {
+ 		if strings.HasPrefix(file.Name(), ".") {
+ 			continue
+ 		}
+ 
+ 		fmt.Println(filepath.Join(dir, file.Name()))
+ 
+ 		if file.IsDir() {
+ 			err := walkDir(filepath.Join(dir, file.Name()))
+ 			if err != nil {
+ 				return err
+ 			}
+ 		}
+ 	}
+ 
+ 	return nil
+ }

func main() {
- 	files, err := ioutil.ReadDir("./")
+ 	err := walkDir("./")
	if err != nil {
		log.Fatal(err)
	}
- 
- 	for _, file := range files {
- 		fmt.Println(file.Name())
- 	}
}
```



### `curl`

### `jq`
