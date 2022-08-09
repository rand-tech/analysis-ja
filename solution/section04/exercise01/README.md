# Exercise1: 整数オーバーフロー

## 解析対象

以下のような整数オーバーフローを発生させるようなコードを見つける。

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("number must be specified")
	}

	n, err := convert(args[0])
	if err != nil {
		return err
	}

	fmt.Println(n, "is positive")

	return nil
}

func convert(numstr string) (int16, error) {
	atoi := strconv.Atoi
	n, err := atoi(numstr)
	if err != nil {
		return 0, err
	}

	if n < 0 {
		return 0, fmt.Errorf("%d is not positive number", n)
	}

	return int16(n), nil // want "integer overflow"
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
