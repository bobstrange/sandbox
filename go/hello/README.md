# How to write Go code

ref: https://golang.org/doc/code.html

前提: `GO111MODULE=on` もしくは `auto` になっていること

## プロジェクトの作成
(example.com/user/hello という module を作る場合)

```shell
mkdir hello && cd hello
go mod init example.com/user/hello
```

Hello, World を出力するプログラム

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

実行ファイルの生成

```shell
go install example.com/user/hello
```

実行ファイルは `$GOPATH/bin/hello` に生成される
(`GO_BIN` が設定されている場合はそちらに)


