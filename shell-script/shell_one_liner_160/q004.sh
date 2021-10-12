#!/bin/bash

# tmp ディレクトリを用意

mkdir -p tmp
cd tmp
ls -U | xargs -P $(nproc) rm 2>/dev/null
seq 100000 | sed 's/^/echo $RANDOM > /' | bash

# Q. ファイル内に 10 という数字が書かれているファイルを削除する

grep -l '^10$' -R | xargs rm

