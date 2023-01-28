# Mastering shell script memo

対応する英語が必要な場合はを参照する
[ref (bash man)](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html)

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

### コマンド置換

[Command Substitution](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Command-Substitution)
[SC2006](https://www.shellcheck.net/wiki/SC2006)

backticks `` より、$() が推奨されている

```bash
# $()
current_dir1=$(pwd)

# ``
current_dir2=`pwd`
```

> 1. It has a series of undefined behaviors related to quoting in POSIX.
> 2. It imposes a custom escaping mode with surprising results.
> 3. It's exceptionally hard to nest.

未定義の振る舞い、想定しづらい escape の mode がある、Nest したときに読みづらい
