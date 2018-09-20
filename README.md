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

POST できるように変更

```diff
-func curl(url string) error {
+func curl(url string, input io.Reader) error {
-	resp, err := http.Get(url)
+	resp, err := http.Post(url, "application/json", input)
 	if err != nil {
 		return err
 	}
 	defer resp.Body.Close()
```

```diff
-	err := curl(url)
+	err := curl(url, os.Stdin)
 	if err != nil {
 		log.Fatal(err)
 	}
```

パイプで POST ボディを渡す

```bash
$ echo '{"foo":"bar"}' | go run *.go https://jsonplaceholder.typicode.com/todos
{
  "foo": "bar",
  "id": 201
}
```


### `jq`

#### ただ JSON をパースするだけ

```diff
+	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
+		log.Fatal(err)
+	} else {
-		err := curl(url, os.Stdin)
+		err := jq(url, input)
 		if err != nil {
 			log.Fatal(err)
 		}
+	}
```

新しいファイル `jq.go`

```go
package main

import (
	"encoding/json"
	"fmt"
)

func jq(path string, input []byte) error {
	var value map[string]string
	err := json.Unmarshal(input, &value)
	if err != nil {
		return err
	}

	fmt.Println(value)
	return nil
}
```

`sample.json`

```json
{
  "foo": "bar"
}
```

```bash
$ cat sample.json | go run *.go dummy
map[foo:bar]
```

#### JSON 文字列に戻す

`sample.json`

```diff
 {
   "foo": "bar",
+  "baz": {
+    "qux": "quux"
+  }
 }
```

`jq.go`

```diff
 func jq(path string, input []byte) error {
-	var value map[string]string
+	var value map[string]interface{}
 	err := json.Unmarshal(input, &value)
 	if err != nil {
 		return err
 	}
 
+	b, err := json.Marshal(value)
+	if err != nil {
+		return err
+	}
+
-	fmt.Println(value)
+	fmt.Println(string(b))
	return nil
 }
```

#### 必要な部分だけ取り出す

```diff
```
