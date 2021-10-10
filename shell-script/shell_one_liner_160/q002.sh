#!/bin/bash

# img ディレクトリ配下にある png 形式の画像を convert コマンドを使用して jpg 形式に変換する

cd -- "${0%/*}"

cd "./shellgei160/qdata/2/img/"

echo "xargs -P 1"
time ls *.png | sed 's/\.png$//' | xargs -I{} convert {}.png {}.jpg


echo "xargs -P 2"
time ls *.png | sed 's/\.png$//' | xargs -P 2 -I{} convert {}.png {}.jpg

echo "xargs -P 3"
time ls *.png | sed 's/\.png$//' | xargs -P 3 -I{} convert {}.png {}.jpg

echo "xargs -P 4"
time ls *.png | sed 's/\.png$//' | xargs -P 4 -I{} convert {}.png {}.jpg

echo "xargs -P 6"
time ls *.png | sed 's/\.png$//' | xargs -P 6 -I{} convert {}.png {}.jpg

echo "xargs -P 8"
time ls *.png | sed 's/\.png$//' | xargs -P 8 -I{} convert {}.png {}.jpg
