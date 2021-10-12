#/bin/bash

cd -- ${0%/*}

# mkdir ./tmp
cd ./tmp
# seq 10000 | xargs -P6 touch

# Q. tmp のファイルの頭に 0 をつけてファイル名を 7 桁に揃える

# ls だと時間がかかるので find を使う
find | head | sed 's/\.\///' | grep -v ^0 | grep -v "^\\.$" | awk '{print($1, sprintf("%07d", $1))}' | xargs -n2 mv

