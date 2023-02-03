#!/bin/bash

read -p "Which file types do you want to backup " file_ext
read -p "Which directory do you want to backup to " dirname

test -d "./${dirname}" || mkdir -m 700 "./${dirname}"

find ./bin -path "./${dirname}" -prune -o \
  -name "*${file_ext}" -exec cp {} "./${dirname}" \;
exit 0
