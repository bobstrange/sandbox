#!/bin/bash

cd -- ${0%/*}
tmp=${pwd}
cd shellgei160/qdata/21

for file in $(echo */*/*); do
  echo "$file"
done

cd ${tmp}
