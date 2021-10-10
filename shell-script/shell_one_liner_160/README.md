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
