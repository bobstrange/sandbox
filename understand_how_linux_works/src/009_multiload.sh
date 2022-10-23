#!/bin/bash

MULTICPU=0
PROGNAME=$0
SCRIPT_DIR=$(cd $(dirname $0); pwd)

usage() {
  exec >&2
  echo "$PROGNAME [-m] <プロセス数>
  所定の時間動作するプロセスを <プロセス数> で指定した数だけ動作させて、全ての終了を待受け、
  各プロセスで動作した時間を出力する。
  デフォルっとでは、1  CPU で動作
  "

  exit 1
}

while getopts "m" OPT; do
  case $OPT in
    m) MULTICPU=1 ;;
    \?) usage ;;
  esac
done

shift $((OPTIND - 1))
if [ $# -lt 1 ]; then
  usage
fi

CONCURRENCY=$1
if [ $MULTICPU -eq 0 ]; then
  # CPU 0 のみで実行
  taskset -p -c 0 $$ > /dev/null
fi

for ((i=0; i<$CONCURRENCY; i++)); do
  time "${SCRIPT_DIR}/008_load.py" &
done

for ((i=0; i<$CONCURRENCY; i++)); do
  wait
done


