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


## package の import

新しく package を作って、それを import する

package 名のディレクトリを作って、そこに実装する。

```shell
mkdir morestrings
touch morestrings/reverse.go
```

./morestrings/reverse.go

```go
package morestrings

func ReverseRunes(s string) string {
    ...
}
```

`ReverseRunes` は大文字から始まっているので、export され、`morestrings` を import している他の package から使用できる。

./hello.go

```go
...
import (
    "fmt"
    "example.com/user/hello/morestrings"
)
...
```

### Remote の module の import

import は、remote の path を書くこともできる。

./hello.go

```go
...
import (
    "fmt"
    "example.com/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)
...
```

`go install` や、 `go build` や、`go run` を実行したときに依存は自動的にダウンロードされて、 `go.mod` ファイルにバージョンとともに記載される

ダウンロードされたものは、`$GOPATH/pkg/mod` に存在する。
