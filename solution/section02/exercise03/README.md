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

## 実行方法

```
$ go build .
$ ./exercise03 testdata/a.go
```
