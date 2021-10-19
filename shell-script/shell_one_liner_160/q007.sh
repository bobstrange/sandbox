#!/bin/bash

cd -- ${0%/*}
tmp="$(pwd)"

cd ./shellgei160/qdata/7

cat kakeibo.txt
cat kakeibo.txt | awk '{print $1,$2,$3 * 1.1}'
cd "${tmp}"
