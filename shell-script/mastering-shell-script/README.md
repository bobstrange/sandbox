# Mastering shell script memo

## Chapter 1

コマンドの種類を判定する `type` command

```bash
❯ type ls
ls is an alias for ls --color=auto
❯ type -a ls
ls is an alias for ls --color=auto
ls is /bin/ls
```

引数

`$0` スクリプトそのもの
`${n}` n 番目の引数
`$#` 引数の数 (`#` number)
`$*` 全ての引数
