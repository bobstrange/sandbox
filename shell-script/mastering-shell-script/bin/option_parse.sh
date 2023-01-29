#!/bin/bash

while [ -n "${1}" ]; do
  case "${1}" in
    -a) echo "-a option used" ;;
    -b) param=${2}
        echo "-b option used with value ${param}" ;;
    -c) echo "-c option used" ;;
    --) shift
        break ;;
  esac
  shift
done

num=1

for param in $@; do
  echo "num: ${num}: ${param}"
  num=$(( num + 1 ))
done
