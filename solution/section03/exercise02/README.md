# Exercise2: 整数オーバーフロー

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
	n, err := strconv.Atoi(numstr)
	if err != nil {
		return 0, err
	}

	if n < 0 {
		return 0, fmt.Errorf("%d is not positive number", n)
	}

	return int16(n), nil
}
```

実行すると以下のようになる。
大きな値を入れた場合におかしくなる。

```
$ go run a.go -1
-1 is not positive number
exit status 1
$ go run a.go 10000
10000 is positive
$ go run a.go 100000
-31072 is positive
```

## テスト方法

`go mod tidy`は初回だけ必要。

```
$ go mod tidy
$ go test
```

## 実行方法

```
$ go build ./cmd/exercise02
$ go vet -vettool=`pwd`/exercise02 testdata/src/a/a.go
```
