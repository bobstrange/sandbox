#!/bin/bash

cd -- ${0%/*}
tmp=$(pwd)

cd shellgei160/qdata/9


cat log_range.log | awk '$4 " " $5 >="[24/Dec/2016 21:00:00]" && $4 " " $5 < "[25/Dec/2016 03:00:00]"'


cd "$tmp"

