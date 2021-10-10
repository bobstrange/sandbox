#!/bin/bash

# ファイル名の一覧を記載した files.txt から .exe の拡張子を持つファイルだけ抜き出す

cd -- "${0%/*}"

FILEPATH=$(find shellgei160 | grep files.txt)

grep "\.exe$" "${FILEPATH}"

# 別解
# cat "${FILEPATH}" | sed -n '/\.exe$/p'
cat "${FILEPATH}" | awk '/\.exe$/'
