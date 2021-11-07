#!/bin/bash

tmp=$(pwd)
cd -- ${0%/*}

n="XYZ"

(
  for i in {A..C}; do
    n+=${i}
    echo ${n}
  done
)

echo ${n}

cd "$tmp"
