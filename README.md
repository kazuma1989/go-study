# Go に触れてみよう

## インストール

- https://golang.org/doc/install#install
- https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go


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
+func walkDir(dir string) error {
+	files, err := ioutil.ReadDir(dir)
+	if err != nil {
+		return err
+	}
+
+	for _, file := range files {
+		if strings.HasPrefix(file.Name(), ".") {
+			continue
+		}
+
+		fmt.Println(filepath.Join(dir, file.Name()))
+
+		if file.IsDir() {
+			err := walkDir(filepath.Join(dir, file.Name()))
+			if err != nil {
+				return err
+			}
+		}
+	}
+
+	return nil
+}
 
 func main() {
-	files, err := ioutil.ReadDir("./")
+	err := walkDir("./")
 	if err != nil {
 		log.Fatal(err)
 	}
-
-	for _, file := range files {
-		fmt.Println(file.Name())
-	}
 }
```


### 新しいコマンドを作るために、ファイル分け

`main.go`

```diff
-func walkDir(dir string) error {
-	files, err := ioutil.ReadDir(dir)
-	if err != nil {
-		return err
-	}
-
-	for _, file := range files {
-		if strings.HasPrefix(file.Name(), ".") {
-			continue
-		}
-
-		fmt.Println(filepath.Join(dir, file.Name()))
-
-		if file.IsDir() {
-			err := walkDir(filepath.Join(dir, file.Name()))
-			if err != nil {
-				return err
-			}
-		}
-	}
-
-	return nil
-}
 
 func main() {
 	err := walkDir("./")
 	if err != nil {
 		log.Fatal(err)
 	}
 }
```

新しいファイル `ls.go`.
処理には変化なし。

```go
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
```

実行は `go run *.go` に変化。


### `curl`

`main.go`

```diff
 package main
 
 import (
 	"log"
+	"os"
 )
 
 func main() {
+	args := os.Args[1:]
+	l := len(args)
+	var url string
+	switch {
+	case l == 0:
+		log.Fatal("Need 1 argument.")
+	case l == 1:
+		url = args[0]
+	case l >= 2:
+		log.Fatal("Too many arguments.")
+	}
+
-	err := walkDir("./")
+	err := curl(url)
 	if err != nil {
 		log.Fatal(err)
 	}
 }
```

新しいファイル `curl.go`

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func curl(url string) error {
	resp, err := http.Get(url)
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
```


### `jq`
