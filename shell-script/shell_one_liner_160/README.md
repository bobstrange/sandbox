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
