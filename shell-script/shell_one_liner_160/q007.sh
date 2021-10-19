#!/bin/bash

cd -- ${0%/*}
tmp="$(pwd)"

cd ./shellgei160/qdata/7

cat kakeibo.txt | \
  awk '{ tax = ($1 < "20191001" || $2 ~ "^*") ? 1.08 : 1.1 ; print $0, tax }' | \
  awk '{ print int($3 * $4) }' | \
  awk '{ total += $1 }END{ print total }'

cd "${tmp}"
