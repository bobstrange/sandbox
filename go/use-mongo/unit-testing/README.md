# unit testing

[ref](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package) の写経

## Table Driven Test

テスト対象の構造体 (入力と欲しい結果) を作る

```go
type example struct {
    arg1, arg2 int
    want int
}
```

テストケースを用意する

```go
var testCases = []example{
    {
        arg1: 10, arg2: 20, want: 30
    },
    ...
}
```

ループ回して実行する

```go
func TestSomething(t *testing.T) {
    for _, testCase := range testCases {
        if output := Something(testCase.arg1, testCase.arg2); output != testCase.want {
            t.Fail("output %q expected %q", output, testCase.want)
        }
    }
}
```

## Coverage

```shell
go test -coverprofile=<出力ファイル> ./...
```

で、テストカバレッジを出力できる


## Show Coverage on html

```shell
go tool cover -html=<↑で出力されたファイル>
```

で、カバレッジをブラウザで表示できる

## Benchmark

BenchmarkXXX で始まる関数を作る
引数として、`*testing.B` を渡す

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(10, 20)
    }
}
```

```shell
go test -bench=. ./
```

ベンチマークを実行する
