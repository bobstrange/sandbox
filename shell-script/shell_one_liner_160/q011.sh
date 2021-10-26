#!/bin/bash

cd -- ${0%/*}
tmp=$(pwd)

cd shellgei160/qdata/11

cat gijiroku.txt | \
  xargs -n2 | \
  sed 's/^すず/鈴木/;s/^さと/佐藤/;s/^やま/山田/' | \
  sed 's/ /: /' | \
  sed 's/$/\n/'

cd "$tmp"
