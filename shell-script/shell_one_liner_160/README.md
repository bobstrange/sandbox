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

