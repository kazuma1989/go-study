# Go に触れてみよう

## インストール

https://golang.org/doc/install#install


## CLI ツールを作る

### `ls`

再帰しない版。

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

再帰版。

```diff
```



### `curl`

### `jq`
