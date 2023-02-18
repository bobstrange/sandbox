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
[ EXPRESSION ] && [ EXPRESSION ]
## or
test EXPRESSION -o EXPRESSION
[ EXPRESSION ] || [ EXPRESSION ]
```

#### 文字列の比較

`==` は `bash` 特有

[参考](https://stackoverflow.com/a/20449556/2571636)

```bash
[ "${USER}" = root ]

# 否定
[ ! "${USER}" = root ]
[ "${USER}" != root ]
! [ "${USER}" = root ]

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

### ファイルの比較

```bash
# ディレクトリかどうか directory
[ -d /home ]

# ファイルが存在するかどうか exists
[ -e /bin/bash]
[ -e /home ]
[ -e /hoge ]

# ファイルが通常のファイルかどうか
[ -f /home ]
[ -f /bin/bash ]
```

## Chapter 5

### 特殊パラメーター

`$-` 設定されているシェルオプションを表示

### デフォルト値

```bash
[ -z "$name" ] && name="default value"

# パラメーター置換 (parameter substitution)
name=${parameter-"default value"}
# ヌル値を持つパラメーターが宣言されている場合
parameter=
${parameter:-default}
```

[ref (parameter substitution)](https://tldp.org/LDP/abs/html/parameter-substitution.html)

### パラメータ展開と引用符

[Shell Parameter Expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)

↓ のようなスクリプトは、パラメータが展開される時にスペースが含まれているため、スペースによって引数が区切られてしまい予期しない動作をする。
(失敗する)

```bash
echo "The file contents" > "my file"
FILE="my file"
[ -f $FILE -a -r $FILE ] && cat $FILE
```

以下のように、 `""` で囲むことで予期した通りの動作になる。

```bash
FILE="my file"
[ -f "${FILE}" -a -r "${FILE}" ] && cat "${FILE}"
```

[Shell Check (SC2086)](https://www.shellcheck.net/wiki/SC2086) でも怒ってくれる

### `[[` を使ったテスト

`sh` では使用できないが `bash` や `zsh` で高度な条件を使ってテストができる

- 空白を含む変数を `""` で囲んでいなくても、パラメータ展開される際に単一の引数として解釈される
- `-a` や `-o` ではなく `&&` や `||` が使用できる

```bash
echo "The file contents" > "my file"
FILE="my file"
[[ -f $FILE && -f $FILE ]] && cat "${FILE}"
```

その他 **パターンマッチング** や

```bash
[[ ${FILE} = *.rb ]] && cp "${FILE}" scripts/
```

**正規表現** が使用ｄけいる

```bash
[[ ${FILE} =~ \.rb$ ]] && cp "${FILE}" scripts/
```

### `(( expression ))` arithmetic expression (算術演算)

```bash
(( a = 1 + 2 ))
let a=1+2
a=$(( 1 + 1 ))
```

`-gt` などを `>` に置き換えられる

```bash
COUNT=2
(( COUNT > 1 )) && echo "Count is greater than 1"
```

## Chapter 6

```bash
for user in alice bob john; do
  echo "Hello ${user}"
done

for user in $(who | cut -f1 -d" "); do
  lsof -u "${user}" -a -c bash | grep cwd
done

for var in one "this is two" three "this is four"; do
  echo "var: ${var}"
done
```

### IFS (Internal field separator)

IFS -> フィールドを区切るための文字
デフォルトはスペースと、タブと改行

```bash
data=$(cat <<EOF
Hello, this is a test
This is the second line
And this is the last one
EOF
)

for var in $(echo ${data}); do
  echo "${var}"
done
Hello,
this
is
a
...
```

初めに見つかった `IFS` に設定された文字がスペースだったので、スペースが区切り文字として扱われた。

```bash
# IFS を改行に変更
IFS=$'\n'
for var in $(echo ${data}); do
  echo "line: ${var}"
done
line:   Hello, this is a test
line:   This is the second line
line:   And this is the last one
```

↓のようにして IFS を初期化できる
`local IFS=$' \t\n'    # normalize IFS`

### C 言語スタイルの for ループ

```bash
for (( v=0; v < 10; v++ )); do
  echo "v: ${v}"
done
```

### ループの制御

```bash
# 結果を result.txt に吐き出す
for (( v=0; v < 5; v++ )); do
  echo "${v}"
done > result.txt

for f in *; do
  # continue で後続を skip して次のループ
  [[ -d "${f}" ]] || continue
  chmod 3777 "${f}"
done

for f in *; do
  [[ -d "${f}" ]] && break
done

echo "Found a directory ${f}"
```

### ファイルからの入力の読み込み

```bash
while read server; do
  ping -c1 ${server} && servers_up="${servers_up} ${server}"
  echo ${server}
done < servers.txt
```

## Chapter 7

関数を作成する方法
`function` キーワードは、移植性のため推奨されていない

```bash
function-name() {
  <code to execute>
}

function <function-name> {
  <code to execute>
}
```

### パラメータ

```bash
print_arg() {
  echo "Input is $1"
}

print_arg 111

print_args() {
  echo "input is $@"

  arr=$@
  echo "arr ${arr[*]}"
}
```

### スコープ

```bash
my_value="outer value"

myfunc() {
  my_value="updated"
}

myfunc
echo "${my_value}"
# => updated

myfunc2() {
  local my_value="updated2"
}

myfunc2

echo "${my_value}"
# => updated (更新されない)
```
