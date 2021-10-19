#!/bin/bash

###########################################
# Q. ↓の出力が得られるワンライナーを作る
#     x
#    x
#   x
#  x
# x
###########################################

seq 5 | sort -nr | awk '{for(i=0;i<$1;i++){ printf " "}; print "x" }'
