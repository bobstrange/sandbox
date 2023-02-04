#!/bin/bash

usage="$(basename $0) <filepath> <search string> <operation>"

if [ ! $# -eq 3 ]; then
    echo "${usage}"
    exit 2
fi

[ ! -f "$1" ] && echo "File $1 does not exist" && exit 3

case "$3" in
  [cC])
    mesg="$1 の中で $2 にマッチする行数を数えます"
    opt="-c"
    ;;
  [pP])
    mesg="$1 の中で $2 にマッチする行数を表示します"
    opt=""
    ;;
  [dD])
    mesg="$1 から $2 にマッチする行を除いて全てを表示します"
    opt="-v"
    ;;
  *)
    echo "$1 $2 $3 を評価できません"
    exit 1;
    ;;
esac

echo "$mesg"
grep $opt "$2" "$1"
exit 0
