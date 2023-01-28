# Mastering shell script memo

## Chapter 1

### コマンドの種類を判定する `type` command

```bash
❯ type ls
ls is an alias for ls --color=auto
❯ type -a ls
ls is an alias for ls --color=auto
ls is /bin/ls
```

### 引数

`$0` スクリプトそのもの
`${n}` n 番目の引数
`$#` 引数の数 (`#` number)
`$*` 全ての引数

### 引用符

`''` 変数展開されない
`""` 変数展開される

### スクリプト名

`$0` スクリプトのパス

```bash
#!/bin/bash

echo "\$0 is ${0}"
```

↑ のような `./bin/script_name.sh` を実行する場合

```bash
❯ ./bin/script_name.sh
$0 is ./bin/script_name.sh

# フルパス
❯ ${PWD}/bin/script_name.sh
$0 is /home/bob/dev/src/github.com/bobstrange/sandbox/shell-script/mastering-shell-script/bin/script_name.sh
```

ファイル名のみ取りたい場合は `basename` コマンドを使う

```bash
#!/bin/bash

echo "\$(basename \$0) is $(basename ${0})"
```

ファイル名が出力される

```bash
❯ ./bin/script_name.sh
$(basename $0) is script_name.sh
```

### 配列

```bash
# 定義
myarr=(foo bar baz qux)

# n 番目の要素
echo ${myarr[1]}

# 全部の要素
echo ${myarr[*]}

# 特定の要素を削除
unset myarr[1]

# 全ての要素を削除
unset myarr
```
