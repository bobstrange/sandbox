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

### Debug

`bash -x`
コマンドが実行されたときのコマンドを表示する
(評価結果なども表示される)

分岐などを追う時に便利

例: bin/debug.sh

```bash
#!/bin/bash

# %w   週の何日目かを表す値 (0..6); 0 を日曜日とする
day_of_the_week=$(date +%w)
date=$(date --iso-8601='seconds')

if [ "${day_of_the_week}" -eq 0 ]; then
  echo "Today is Sunday: ${date}"
else
  echo "Today is not Sunday: ${date}"
fi
```

```bash
❯ bash -x ./bin/debug.sh
++ date +%w
+ day_of_the_week=0
++ date --iso-8601=seconds
+ date=2023-01-29T01:08:50+09:00
+ '[' 0 -eq 0 ']'
+ echo 'Today is Sunday: 2023-01-29T01:08:50+09:00'
Today is Sunday: 2023-01-29T01:08:50+09:00
```

### VSCode Bash Debug

- VSCode の Bash Debug Extension をインストール
- コマンドパレットで `Debug:AddConfiguration` -> `Bash Debug` を選択し、設定を作る
  - .vscode/launch.json に debug 用の設定が追加される

```jsonc
{
 "version": "0.2.0",
 "configurations": [
  {
    "type": "bashdb",
    "request": "launch",
    "name": "Bash-Debug (simplest configuration)",
    "program": "${file}",
    "cwd": "${fileDirname}"
  }
 ]
}
```

`cwd` に、`${fileDirname}` を設定しておくことで、shellscript を開いた状態で `F5` で debug ができる

## Chapter 2

### echo コマンド

デフォルトの挙動は、引数 + 改行を出力

```bash
# 標準
❯ echo "Example text"
Example text

# -n で改行が出力されない
❯ echo -n "Example text"
Example text%

# -e + \c (これ以上の文字を出力しない)
❯ echo -e "Example text\c aaaaa"
Example text%
```

### read コマンド

`read` コマンドはシェル組み込みコマンド

`-n` で入力文字数を制限
※ zsh では `-k`
`-p` でプロンプトの文字を設定
※ zsh では `変数名?プロンプト`

[ref(zsh shell builtin command)](https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html)

```bash
read -p "Please enter your name: " name
echo "Hi ${name}!"

read -n1 -p "Press any key to exit"
echo "Bye!"
```

入力の可視性を制御することもできる

```bash
read -s -p "Please enter password:" password
```

### オプション

```bash

while [ -n "${1}" ]; do
  case "${1}" in
    -a) echo "-a option used" ;;
    # 引数が値を受け取るケース
    -b) param=${2}
        echo "-b option used with value ${param}" ;;
    -c) echo "-c option used" ;;
    # -- 以降はオプションをチェックしない
    --) shift
        break ;;
  esac
  shift
done

num=1

for param in $@; do
  echo "num: ${num}: ${param}"
  num=$(( num + 1 ))
done

```

```bash
❯ bin/option_parse.sh -a -b p1 -c -- p1 p3
-a option used
-b option used with value p1
-c option used
num: 1: p1
num: 2: p3
```

[Shell builtin commands](https://www.gnu.org/software/bash/manual/bash.html#Shell-Builtin-Commands)

> Unless otherwise noted, each builtin command documented as accepting options preceded by ‘-’ accepts ‘--’ to signify the end of the options. The :, true, false, and test/[ builtins do not accept options and do not treat ‘--’ specially. The exit, logout, return, break, continue, let, and shift builtins accept and process arguments beginning with ‘-’ without requiring ‘--’. Other builtins that accept arguments but are not specified as accepting options interpret arguments beginning with ‘-’ as invalid options and require ‘--’ to prevent this interpretation.

`--` (double dash) を渡すことで、オプションがこれ以降受け付けないということを示す。

### オプションを解析するコマンド getopt / getopts

[ref(シェルスクリプト オプション解析 徹底解説 (getopt / getopts))](https://qiita.com/ko1nksm/items/cea7e7cfdc9e25432bab#getopt-%E3%81%A8-getopts-%E3%81%AE%E9%81%95%E3%81%84)

基本的には、→の方が高機能
「getopt（オリジナル版）」<「getopts」<「getopt（GNU版）」

## Chapter 3

### 実行パス

最後に実行したコマンドの exit code を表示する

```bash
echo $?
```

例: 特定のディレクトリにいるかチェックして、移動する

```bash
test ${PWD} == ${HOME} || cd ${HOME}
```

### test command

```bash
test EXPRESSION

# 真偽の反転 (※ ! と EXPRESSION の間に半角スペースが必要)
test ! EXPRESSION

# 複数条件
## and
test EXPRESSION -a EXPRESSION
## or
test EXPRESSION -o EXPRESSION
```

#### 文字列の比較

`==` は `bash` 特有

[参考](https://stackoverflow.com/a/20449556/2571636)

```bash
[ "${USER}" = root ]
[ ! "${USER}" = root ]

# 文字列長が 0 より大きいか?
[ -n "something" ] #=> 0
[ -n "" ]          #=> 1

# 文字列長が 0 か
[ -z "" ]          #=> 0
[ -z "something" ] #=> 1
```

### 整数の比較

```bash
# equal to
[ num1 -eq num2 ]
# not equal to
[ num1 -ne num2 ]
# greater than or equal to
[ num1 -ge num2 ]
# greater than
[ num1 -gt num2 ]
# less than or equal to
[ num1 -le num2 ]
# less than
[ num1 -lt num2 ]
```
