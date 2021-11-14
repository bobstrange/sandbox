# シェル・ワンライナー 160 本ノックのメモ

## 計算

```bash
echo '1+1' | <something>
```

`<something>` の部分に何が入るか？

```bash
echo '1+1' | bc

# Ruby なら $stdin.read が使える
echo '1+1' | ruby -e 'puts eval($stdin.read())'
echo '1+1' | ruby -e 'puts $stdin.read().split("+").map(&:to_i).sum'

# 標準入力をそのままコードとしても評価できる
echo 'puts 1+1' | ruby
# つまり ↓ができる
echo '1+1' | sed 's/^/puts /' | ruby

echo $(( 1 + 1 ))
```

## ファイルへの保存

```bash
echo '1+1' | bc > result.txt
```

## ファイルとディレクトリの操作

`nautilus` を実行すると、ファイラが起動する
(ファイラの名前 `nautilus` だったのね、、、)

## ファイルのパーミッション

chmod `+` で追加 `-` で削除
`rwx` が権限 `read` `write` `execute`

```bash
echo '1+1' | bc > a
chmod -r a
cat a
chmod +r a
cat a
rm a
```

## sed

```bash
# Q クロロメチルエチルエーテル
echo クロロエチルエチルエーテル | sed 's/エ/メ/'

# Q2 クロロエチルメチルエーテル
echo クロロエチルエチルエーテル | sed 's/エチルエ/エチルメ/'

# Q3 クロロエチルエチルエーテル
echo クロロメチルメチルエーテル | sed 's/メ/エ/g'

# Q4 クロロエチルエチルエーテル
echo クロロエチルエーテル | sed 's/エチル/&&/'

# Q5 クロロエチルメチルエーテル
echo クロロメチルエチルエーテル | sed -E 's/(メチル)(エチル)/\2\1/'
```

### Q4 マッチした部分の再利用(後方参照)

`&` でマッチした箇所の再利用ができる

### Q5 マッチした部分の再利用その2

拡張正規表現を利用すると、マッチした部分の一部を置き換えられる

## grep

```bash
# 0 を含むもの
seq 100 | grep "0" | xargs

# 8 から始まるもの
seq 100 | grep "^8" | xargs

# 8 で終わるもの
seq 100 | grep "8$" | xargs

# 8任意の文字列 にマッチするもの
seq 100 | grep "8." | xargs

# 1 で始まり 0 が任意の数で行末になるもの
seq 100 | grep "^10*$" | xargs

# 偶数
seq 100 | grep "[02468]$" | xargs

# 奇数
seq 100 | grep "[^02468]$" | xargs

# 同じ文字が 2 文字続く
seq 100 | grep -E "^(.)\1$" | xargs
```

```bash
echo 中村 山田 田代 上田 | grep -o "[^ ]田" | xargs

echo 中村 山田 田代 上田 | sed 's/[^ ][^ 田]//g'
echo 中村 山田 田代 上田 | tr ' ' '\n' | grep "田$" | xargs
```

`/[^ ][^ 田]//` スペース以外, スペースと田以外 にマッチングした場合に消す
`tr ' ' '\n'` で、半角スペースを、改行に置き換えてから 田 で終わる文字にマッチング


## awk

`awk` は `grep` にプログラム機能をつけたものと考えると良いらしい


```bash
# grep のように正規表現でマッチング
seq 5 | awk '/[2 4]/'

# $1 は読み込んだ行の 1 列目の文字列
seq 5 | awk '$1 % 2 == 0'
```

`print` を使う

```bash
seq 5 | awk '$1 % 2 == 0{printf("%s 偶数\n", $1)}'
seq 5 | awk '$1 % 2 == 0{print($1, "偶数")}'
seq 5 | awk '$1 % 2 == 0{print $1, "偶数" }'
```

複数の処理を書く場合

※ 条件 -> **パターン** 処理 -> **アクション** パターンとアクションの組 -> **ルール**

```bash
seq 5 | awk '$1 % 2 == 0{print $1, "偶数"}$1 % 2{print $1, "奇数"}'
seq 5 | awk '$1 % 2 == 0{print $1, "偶数"}$1 % 2 == 1{print $1, "奇数"}'
```

`BEGIN` `END` を使う

```bash
seq 5 | awk 'BEGIN{a=0}$1 % 2 == 0{print $1, "偶数"}$1 % 2{print $1, "奇数"}{a += $1}END{print "合計", a}'
```

ルールの間に区切り文字が無いのが辛い :-(

`BEGIN` パターン -> 1 行目の処理の前
`END` パターン -> 最終行の処理の後

`{a += $1}` アクションは、パターンが無いので、全部の行に対して適用される
awk では変数定義が不要なので、この例だと `BEGIN` パターンは省略できる

おすすめされている参考文献は 「シェル芸」に効く AWK 処方箋らしい


## sort と uniq

```bash
seq 5 | awk '{ print $1 % 2 ? "奇数" : "偶数" }' | sort | uniq -c | awk '{ print $2,$1 }'
```

`uniq` は重複行を消すコマンドだが、重複は連続した行である必要がある (つまり sort 済みでなければならない)
`-c` オプションで、重複の件数を出力できる

`sort` のオプションとして、 `-n` 数字順でソート `-k` で列を指定してソートをよく使用する

出力結果を数字順にソートする場合

```bash
seq 5 | awk '{ print $1 % 2 ? "奇数" : "偶数" }' | sort | uniq -c | awk '{ print $2,$1 }' | sort -k 2,2n
```

`-k 2,2` で 2 列目から 2 列目を基準にソートするという意味
`n` は対象の列に `-n` オプションを適用するという意味

`-n` をつけないと結果が変わる場合もある

```bash
seq 19 | awk '{ print $1 % 2 ? "奇数" : "偶数" }' | sort | uniq -c | awk '{ print $2,$1 }' | sort -k 2,2n
偶数 9
奇数 10

# 1 と 9 の比較で 1 が先になる
seq 19 | awk '{ print $1 % 2 ? "奇数" : "偶数" }' | sort | uniq -c | awk '{ print $2,$1 }' | sort -k 2,2
奇数 10
偶数 9
```

awk を使用するだけでも

```bash
seq 19 | awk '{ print $1 % 2 ? "奇数" : "偶数"}' | awk '{a[$1]++}END{for(k in a)print k, a[k]}'
```

`a[$1]` で、連想配列
`for(k in obj)` で、obj のキーについてループできる
`for(i=1;i<5;i++)` 的なループもできる

## xargs

`xargs` は、コマンドに引数を渡して実行してもらうためのコマンド

```bash
# 1 ~ 4 までのディレクトリを作成
seq 4 | xargs mkdir

# 1 ~ 4 のディレクトリを削除
seq 4 | xargs rmdir

# 1, 3 というディレクトリを 2, 4 に変更
mkdir 1 3
seq 4 | xargs -n2 mv

# dir_1 ~ dir_4 を作成
seq 4 | xargs -I{} mkdir dir_{}
```

`xargs -n <num>` で、入力された文字列を指定した個数ずつコマンドに渡すという意味になる

```bash
seq 5 | xargs -n 2
1 2
3 4
5
seq 5 | xargs -n 3
1 2 3
4 5
```

`xargs -I{}`
`xargs` が受け取った文字列で `-I` に指定したパターンを置き換える

```bash
seq 4 | xargs -I{} echo dir_{}
dir_1
dir_2
dir_3
dir_4
```

## bash によるメタプログラミング

`bash` はコマンドを並べた命令をパイプから受け取ることができる

```bash
seq 4 | awk '{print "mkdir " ($1 % 2 ? "odd_": "even_") $1}'
mkdir odd_1
mkdir even_2
mkdir odd_3
mkdir even_4

# bash に渡すことで、mkdir が実行される
seq 4 | awk '{print "mkdir " ($1 % 2 ? "odd_": "even_") $1}' | bash
```

※ 三項演算子の部分を `()` で囲む

## ファイルの操作

`find` 存在するディレクトリが列挙される

```bash
find shellgei160 | grep files.txt
```

## シェルの基本

### 標準入出力

コマンドにファイルの中身を入力する時

`cat` + `|`:  `cat file | wc -l`
入力のリダイレクト記号 `<`: `wc -l < file`

多くのコマンドが、コマンドが別のコマンドに渡すべきデータ(処理結果)と、そうでないデータ(ヘルプメッセージ)を区別している
->  標準出力 と 標準エラー出力

標準エラー出力をパイプで渡すには

```bash
sed 2>&1 | wc -l
sed |& wc -l
sed |& less
```

`n>&m` は n 番の出力を m 番に向ける
`|&` は標準出力、標準エラー出力をパイプで渡す

### 文字列連結と置換

パラメータの展開 (Parameter Expansion) (`man bash` の Parameter Expansion を参考に)

```bash
a="私は"
b="俳優よ"
c="${a}${b}"; echo ${c}
a+="${b}"; echo ${a}
b="${a:0:1}${a:2:2}"; echo ${b}
c=${a/俳優/排骨麺}; echo ${c}
```

#### Parameter Expansion (変数展開)

部分文字列

- `${parameter:offset}`
  - `foo="test"; echo ${foo:1}` => `est`
`${parameter:offset:length}`
  - `foo="testdata"; echo ${foo:1:3}` => `est`

長さ

- `${#parameter}
  - `foo="test"; echo ${#foo}` => 4

特定のパターンを削除
Remove matching prefix pattern

- `${parameter#word}
  - `#` だと 最短一致
- `${parameter##word}
  - `##` だと 最短一致

```bash
foo="foo/bar/baz.json.gz"; echo ${foo#*/}
bar/baz.json.gz
foo="foo/bar/baz.json.gz"; echo ${foo##*/}
baz.json.gz
```

Remove matching suffix pattern

- ${parameter%word}
  - `%` だと 最短一致
- ${parameter%%word}
  - `%%` だと 最短一致

```bash
foo="foo/bar/baz.json.gz"; echo ${foo%/*}
foo/bar
foo="foo/bar/baz.json.gz"; echo ${foo%%/*}
foo
```

### 算術式展開

`$(())` の中に計算式を入れると計算できる (`man` の `ARITHMETIC EVALUATION`)

```bash
a=6
b=2
echo $((a + b)) $((b - a)) $((a * b)) $((a / b)) $((b << a))
```

ビット演算もできる
`2 << 6` は 6 ビットシフト
`10` -> `10000000` -> 2^8 -> 128

### クォートと変数

`'` シングルクォート -> 囲われた文字列について Bash は展開などを何も行わない
`"` ダブルクォート -> 変数や、特殊記号などが解釈される

空文字があることで、スクリプトが想定外の挙動になってしまう場合があるので、変数を参照するときには、ダブルクォートで囲んで置くと良い

```bash
a=("$SHELL" "$LANG" "$USER")
declare -A b
b["SHELL"]="$SHELL"
b["LANG"]="$LANG"
b["USER"]="$USER"

echo ${a[2]}
echo ${b["LANG"]}

# 要素を全て表示
echo ${a[@]}
echo ${b[@]}

# 要素数を表示
echo ${#a[@]}
echo ${#b[@]}

# キーを表示
echo ${!b[@]}
```

### プロセス

#### バックグラウンドジョブ

`&`: バックグラウンドでジョブを実行するための記号
ジョブ: 1 個以上のコマンドが組み合わさった処理をシェルが管理するための単位

例: 5 つの sleep が 1 つのジョブとして扱われる

`[1]` は ジョブ ID

```bash
 sleep 100 | sleep 100 | sleep 100 | sleep 100 | sleep 100 &
[1] 29479
```

`fg <job_id>` でバックグラウンドジョブをフォアグラウンドに持ってくることができる
フォアグラウンドのジョブを `ctrl + z` で停止して、 `bg` でバックグラウンドで再開できる

#### プロセスの親子関係

`pstree` コマンドで、プロセスの親子関係を表示することができる

```bash
pstree -T | head -n 10
```

sleep を実行してから、`pstree` で確かめてみる

```bash
 sleep 100 | sleep 100 | sleep 100 | sleep 100 | sleep 100 &
```

```bash
pstree -T | grep -A5 bash
        |         |      |             `-zsh---bash-+-grep
        |         |      |                          |-pstree
        |         |      |                          `-5*[sleep]
        |         |      `-code-+-code---7*[code]
        |         |             `-code
        |         |-dbus-daemon
```

`ps --forest` でも `ps` の出力に親子関係を加えられる

#### ビルトインコマンドと外部コマンド

ファイルの実体があるコマンド (`/bin/bash` など) を外部コマンド
実体が無いコマンド (`cd`, `set`, `read` ) などは、ビルトインコマンド(組み込みコマンド)と呼ぶ
ビルトインコマンドは、シェルの機能として、シェルにtyokusetu puroguramusareteiru .

いくつかのコマンド (`echo` など) は、シェルに直接プログラムされている
ビルトインコマンドのメリットは、高速であること

例: ビルトインコマンドの echo と 外部コマンドの echo

```bash
time for i in {1..1000}; do /bin/echo "$i" > /dev/null; done

real    0m0.762s
user    0m0.573s
sys     0m0.227s

time for i in {1..1000}; do echo "$i" > /dev/null; done

real    0m0.009s
user    0m0.004s
sys     0m0.004s
```

100 倍くらい違う
外部コマンドは、別プロセスを起動するが、ビルトインコマンドは関数実行のようなものなので、呼び出しコストが全然違う

#### ビルトインコマンドと外部コマンドの見分け方

`which` や、 `type`、 `command -v` などのコマンドで、見分けることができる

```bash
which echo
/bin/echo

type echo
echo はシェル組み込み関数です

 command -v echo
echo

which bash
/bin/bash

type bash
bash は /bin/bash です

command -v bash
/bin/bash
```

### サブシェルを使う

例えば `/etc` の中身を確認してから元のディレクトリに戻る場合 ↓のように記載するが

```bash
cd /etc
ls
cd -
```

サブシェルを使用すれば、現在のディレクトリで、書くことができる

```bash
(cd /etc/; ls)
```

```bash
a="きたうらわ"
echo "${a}を逆さにすると$(echo ${a} | rev)"

cat <(echo $a) <(echo を逆さにすると) <(echo $a | rev)
```

`<()` プロセス置換するための書き方
`<()` がファイル名のように扱われ、`()` 内のコマンドの出力が、ファイルの内容のように扱われる

`cat <(echo $a)` で きたうらわ
`cat <(echo hoge) <(echo fuga)` で複数のファイル(プロセス置換されたもの)の内容を表示

hoge
fuga




---

## Q.001

[answer](./q001.sh)

`sed -n` で各行を自動的に出力しない
`/regexp/p` で正規表現にマッチする行のみ出力

## Q.002

[answer](./q002.sh)

`time` でコマンドの実行時間を計測できる

```bash
./q002.sh

real    0m1.939s
user    0m9.813s
sys     0m0.837s
```

`xargs -P` で並列実行ができる

```
xargs -P 1

real    0m1.761s
user    0m9.184s
sys     0m0.790s
xargs -P 2

real    0m1.201s
user    0m9.436s
sys     0m0.789s
xargs -P 3

real    0m1.113s
user    0m9.697s
sys     0m0.763s
xargs -P 4

real    0m1.119s
user    0m10.106s
sys     0m0.593s
xargs -P 6

real    0m1.064s
user    0m9.919s
sys     0m0.641s
xargs -P 8

real    0m1.027s
user    0m9.944s
sys     0m0.559s
```

ページキャッシュが効いているので、速度は線形にはならない

`xargs -P $(nproc)`
`nproc` で使用できるプロセッサの数がわかるので、`-P` に渡せる

## Q.003

[answer](./q003.sh)

ファイル名変更 `rename` コマンドが使えるらしい
(そんなこまんどあるんですね)

`sudo apt install rename` でインストール

`ls` は、ファイルの一覧を出力する時にソート処理を行うので、ファイル名が遅いと処理に時間がかかる
-> `ls -U` オプションを使用する

```bash
ls -U | xargs -P $(nproc) rename 's/^/0000000/;s/0*([0-9]{7})/$1/'
```

`rename` に 2 つのルールを渡している

- `s/^/0000000/` -> ファイル名の先頭に 0000000 を一律つけている
- `s/0*([0-9]{7})/$1/` `([0-9]{7})` で 7桁にマッチして、後方参照している

## Q.004

[answer](./q004.sh)

`grep -R` ディレクトリの内部のファイルを再帰的に読み込む
`grep -l` ファイル名のみ出力

## Q.005

[answer](./q005.sh)

`grep` ではなく `awk` で `pool` で始まる行を抽出できる

```bash
cat shellgei160/qdata/5/ntp.conf| awk '$1 == "pool"{ print $2 }'
```

## Q.006

[answer](./q006.sh)

```bash
seq 5 | awk '{ for(i = 0;i < $1; i++) { printf " "}; print "x" }' | tac
```

`tac` コマンドで出力をひっくり返すことができるらしい

`seq` コマンドは、↓のように使えるので、`FIRST` に 5 を `INCREMENT` に -1 を LAST に `1` を設定し、5 ~ 1 を出力できる

seq [OPTION]... LAST
seq [OPTION]... FIRST LAST
seq [OPTION]... FIRST INCREMENT LAST

```bash
seq 5 -1 1
5
4
3
2
1
```

## Q.007

awk で行全体を表すのは `$0`

```bash
cat ./shellgei160/qdata/7/kakeibo.txt | awk '{ tax = 1.1; print($0, tax) }'
```

$1 が "20191001" 以前もしくは、
`*` がついている行のみ tax を 1.08 にする

`||` or 演算子と、三項演算子、 `~` 正規表現が使用できる

```bash
cat ./shellgei160/qdata/7/kakeibo.txt | awk '{ tax = ( $1 < "20191001" || $2 ~ /^\*/ ); print($0, tax) }'
```

`int()` で切り捨てができる

## q009

`awk '$4 " " $5'` で、4 列目と空白１つと5列目をそのまま連結した文字列になる

```bash
echo "sample1 sample2 sample3" | awk -F ' ' '$1'
```

## q010

[answer](./q010.sh)

`sed` で変換すれば良い

拡張正規表現を使用するので `-E` オプションをつける (もしくは `-r`)
後方参照を使用して、`/## (.*)/\1\n---` と置き換える
`#` の後のスペースは 1 以上の任意なので、 `/## +(.*)/` とした方が良い

## q011

[answer](./q011.sh)

- まず、名前 と 発言を同じ行にするため `xargs -n 2` で、2 行を 1行として扱うようにする
- 次に、`sed` で各種変換を行う

## q012

[answer](./q012.sh)

OR 演算子を使って、 if 文より短く書くことができる

```bash
[ "$1" == "" ] && read num || num = "$1"
echo $(( num * 2 ))
```

`:-` でデフォルト値を指定するやり方 + コマンド置換でもできる

```bash
num="${1:-$(cat)}
echo $((num * 2))
```

## q013

> touch はファイルの時刻関連の記録であるタイムスタンプを編集するコマンド

空ファイルを作るコマンドだと思っていた、、、
確かに、 `touch` という名前からは、タイムスタンプ更新を類推できる、、、

`cat` で存在しないファイルを指定するとエラーになってしまうが、 `touch` で事前に空ファイルを作成しておくことでエラーを回避できるようになる。
ただ、ファイルが存在する場合に、タイムスタンプが更新されてしまうので、ファイルが存在しない場合だけ `touch` を実行したい。

```bash
[ -e unfile ] || touch unfile
```

`[` ( `test` ) コマンドの `-e` オプションで、ファイルの有無をチェックできる

## q14

ループを回すやり方を覚えるための問題

カウンタ `n` を使う場合

```bash
n=1
while [ $n -le 100 ]; do
  n=$(( n + 1 ))
done
```

`seq` で数字を作って while を使う場合

```bash
seq 100 | while read n; do ... done
```

for で回す場合

```bash
for n in $(seq 100); do ... done
for n in {1..100}; do ... done
for (( n = 1; n <= 100; n++ )); do ... done
```

`xargs` + `bash` で組み合わせる場合

```bash
xargs 100 | -I@ bash -c " ... ; sleep 1"
```

## q015

`sed` 使って解いたけど `sed` ダメなのか :-(

`\U` と `\L` で uppercase と lower case

`tr` (translate) でも良い `echo something | tr '[:lower:]' '[:upper:]'

bash だけでなんとかする場合は、パイプから変数に文字を取り込む必要があるらしい

- q014 でもやった、`while read var; do ... done` で、変数を読み取れる
- `echo something | bash -c 'read a; echo $a'` でも読み取れる
- サブシェルでも良い `echo something | (read a; echo $a)`

変数に取り込んだら、bash の変数展開 `${変数名^^}` で大文字にできる

## q016

for 文の処理の時に、サブシェルで処理するようにすることで、
元のシェルの変数が書き換わらないようにする。

単純に `()` で囲えば良いだけ

for 文にパイプをつなぐことで、 for 文がサブシェルで実行される
パイプは、異なるプロセス間でのデータの受け渡し方法 -> 両側のコマンドは別のプロセス

シェルスクリプトで、想定通りに変数が書き換わらない or 書き換わってしまったという問題の元になることが多い

## q017

```bash
echo $(< /etc/passwd)
```

`$(< ファイル)` はコマンド置換みたいなもの
ファイルの中身を引数に書き換えられる

`read` はシェル組み込み関数

```bash
type read
read はシェル組み込み関数です
```

```bash
while read line; do
  echo ${line}
done < /etc/passwd
```

として、ループ全体に対してリダイレクトをすることができる

## 2.3 ブレース展開とファイルグロブ

### ブレース展開

`{1..100}` や、 `{A..C}` -> シーケンス式

```bash
echo {山,上}田
山田 上田

echo {1..5}.{txt,bash}
1.txt 1.bash 2.txt 2.bash 3.txt 3.bash 4.txt 4.bash 5.txt 5.bas

echo {2..10..2}.{txt,bash}
2.txt 2.bash 4.txt 4.bash 6.txt 6.bash 8.txt 8.bash 10.txt 10.bash

echo {山,上}{田,　}
山田 山　 上田 上　
```

`{2..10..2}` というシーケンス式で、1 つおきに数値を出力できる

### ワイルドカードとファイルグロブ

`*` ワイルドカード
※ 正規表現の `*` とはちょっと異なる

`?` 任意の一文字
`[]` 内側の文字列のいずれか

基本的には、正規表現と対応が取れている

`*.png` のような文字列をグロブと呼ぶ

```bash
# 前提
mkdir tmp
cd tmp
touch {1..100}.{txt,bash}
rm 5.txt 25.bash

ls {1..9}.txt
ls: '5.txt' にアクセスできません: そのようなファイルやディレクトリはありません
1.txt  2.txt  3.txt  4.txt  6.txt  7.txt  8.txt  9.txt

# 5.txt が存在しないので、ブレース展開で `ls {1..9}.txt` とやるとエラーがでてしまう

ls ?.txt
1.txt  2.txt  3.txt  4.txt  6.txt  7.txt  8.txt  9.txt

# ブレース展開でもできるが、、、
ls {1,2,6}5.*
15.bash  15.txt  25.txt  65.bash  65.txt

# グロブを使用した方が良い
ls [126]5.*

# grep ではなく、、、
ls ?.* | grep -v -E [29]

# ファイルグロブを使う
# [1345678] と書き下すより [^29] とした方が良い
# [13-8] でも表示できる
ls [^29].*

1.bash  1.txt  3.bash  3.txt  4.bash  4.txt  5.bash  6.bash  6.txt  7.bash  7.txt  8.bash  8.txt```
```

## q020

シェルの文字列展開でもできる

```bash
for file in /usr/*; do
  echo ${file}
done

/usr/bin
/usr/games
/usr/include
/usr/lib
/usr/libexec
/usr/local
/usr/sbin
/usr/share
/usr/src

for file in /usr/*; do
  echo ${file##*/}
done
```

`${変数名##*/}` 除去したい文字列

