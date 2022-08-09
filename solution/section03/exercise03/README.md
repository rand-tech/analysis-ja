# Exercise3: エラーの無視

## 解析対象

以下のようなエラーを無視しているコードを見つける。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi("100")
	fmt.Println(n)
}
```

## テスト方法

`go mod tidy`は初回だけ必要。

```
$ go mod tidy
$ go test
```

## 実行方法

```
$ go build ./cmd/exercise03
$ go vet -vettool=`pwd`/exercise03 testdata/src/a/a.go
```
