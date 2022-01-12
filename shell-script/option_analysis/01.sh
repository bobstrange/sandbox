#!/bin/bash

for OPT in "$@"; do
  case ${OPT} in
    -q | --queue)
      queue=$2
      shift 2
      ;;
    -t | --timeout)
      timeout=$2
      shift 2
      ;;
  esac
done

echo ${queue}
echo ${timeout}
