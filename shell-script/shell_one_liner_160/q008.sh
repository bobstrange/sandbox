#!/bin/bash


cd -- ${0%/*}
tmp=$(pwd)

cd shellgei160/qdata/8

cat access.log | \
  sed -E 's/.*\[(.*)]$/\1/' | \
  awk -F ':' '{ print $2 }' | \
  awk '$1 < 12 { print "am" } $1 >= 12 { print "pm" }' | sort | uniq -c

cd "$tmp"
