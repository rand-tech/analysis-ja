# Section 4: 静的単一代入形式を用いた解析

* [Exercise1: 整数オーバーフロー](./exercise01)

## skeletonの実行方法

-kindでssaを指定する。

```
$ skeleton -kind ssa github.com/gohandson/analysis-ja/solution/section04/exercise01
```

## TODO

以下のコードを検出するようなLinterを静的単一代入形式を用いて作成せよ。

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

* まずは`*ssa.Convert`を探す
* 変換する型が`types.Typ[types.Int16]`または`types.Typ[types.Int32]`か確認
* 変換する値が`*ssa.Extract`であることを確認し、Indexが0か確認する
* ssa.ExtractのTuppleが*ssa.Callの場合、呼び出している関数がstrconv.Atoiか確認する
* 呼び出している関数はssa.CallのCallフィールドからValueを取得し、Objectメソッドでtypes.Objectを取得すれば比較ができる
