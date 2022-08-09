# Exercise3: パッケージ変数の初期化における不正なコマンドの呼び出し

## 解析対象

以下のようなパッケージ変数の初期化で不正なコマンドの呼び出しを行っている箇所を検出する。

```go
package lib

import (
	"os"
	"os/exec"
)

var Message = "hello"

var _ = func() struct{} {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	return struct{}{}
}()
```

## 実行方法

```
$ go build .
$ ./exercise03 testdata/lib/lib.go
```
