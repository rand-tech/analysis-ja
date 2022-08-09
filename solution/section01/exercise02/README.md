# Exercise2: init関数における不正なコマンドの呼び出し

## 解析対象

以下のような`init`関数内で不正なコマンドの呼び出しを行っている箇所を検出する。

```go
package lib

import (
	"os"
	"os/exec"
)

var Message = "hello"

func init() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
```

## 実行方法

```
$ go build .
$ ./exercise02 testdata/lib/lib.go
```
