#!/bin/bash

cd -- ${0%/*}
tmp=$(pwd)

n="XYZ"

(
  for i in {A..C}; do
    n+=${i}
    echo ${n}
  done
)

echo ${n}

cd "$tmp"
