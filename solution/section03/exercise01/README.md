# Exercise1: unsafeパッケージの利用

## 解析対象

以下のような`unsafe`パッケージを利用している箇所を探す。

```go
package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	X [2]string
	Y string
}

func main() {
	t := T{
		X: [...]string{"A", "B"},
		Y: "C",
	}

	xp := uintptr(unsafe.Pointer(&t.X))
	yp := (*string)(unsafe.Pointer(xp + unsafe.Sizeof("")*2))
	fmt.Println(*yp)
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
$ go build ./cmd/exercise01
$ go vet -vettool=`pwd`/exercise01 testdata/src/a/a.go
```
